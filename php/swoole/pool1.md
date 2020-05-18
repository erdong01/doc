### swoole 多进程处理数据
在业务不断变化开发中，难免会碰到需要做数据处理，当数据超过10万+耗时就会变长，这时候需要开多任务处理。

### swoole文档:

https://wiki.swoole.com/#/memory/lock

https://wiki.swoole.com/#/memory/table
####环境：
php 7.4

swoole 4.5

mysql 5.7

#### laravel框架或者基于PDO 需要修改options，config配置需要修改  
```php
[
//\PDO::ATTR_PERSISTENT => true, // 关闭持久连接
\PDO::ATTR_EMULATE_PREPARES => true //开启预模拟预处理语句
];
```

#### 用的是laravel框架不能使用协程，所以只能使用多进程处理数据。
```php

<?php

namespace App\Console\Commands\Common;

use App\Models\Question;
use Illuminate\Console\Command;
use Illuminate\Database\Eloquent\Builder;
use Swoole\Process\Pool;
use Swoole\Table;

class DataTreating extends Command
{
    private $workerNum;//worker进程数
    private $limit;//每一次处理数据条数

    public function handle()
    {
        $starttime = explode(' ', microtime());
        
        ini_set('memory_limit', '8000M');//按需设置内存
        ini_set('always_populate_raw_post_data', '-1');
        $this->limit = 500;
        $this->workerNum = 24;// worker进程数

        $questionSql = $this->questionStl();
        $questionCount = $questionSql->count();
        $cycleIndex = ceil($questionCount / $this->limit);
        $this->worker($cycleIndex);


        $endtime = explode(' ', microtime());
        $thistime = $endtime[0] + $endtime[1] - ($starttime[0] + $starttime[1]);
        $thistime = round($thistime, 3);
        echo "执行耗时：" . $thistime . " 秒。" . PHP_EOL;
    }

    private function worker($cycleIndex)
    {
        $table = new Table(100);
        $table->column('index', Table::TYPE_INT, 8);//数据页数
        $table->column('task_sum', Table::TYPE_INT, 4);
        $table->column('worker_status', Table::TYPE_INT, 4);
        $table->create();
        $table->set("queue", ["index" => 0, "task_sum" => $this->workerNum]);
        if ($cycleIndex < $this->workerNum) {
            $this->workerNum = $cycleIndex;
        }
        $pool = new Pool($this->workerNum);
        $pool->on("WorkerStart", function (Pool $pool, $workerId) use ($table, $cycleIndex) {
            //页数 原子自增
            $table->incr("queue", "index", $incrby = 1);
            $queue = $table->get("queue");
            if ($queue['index'] > $cycleIndex) {
                sleep(100); //超过数据处理页数，直接让进程睡眠。
                return;
            }
            //取出worker进程 原子自减
            $table->decr("queue", "task_sum", $incrby = 1);
            $table->set($workerId, ['worker_status' => 1]);
            \DB::disconnect(); //laravel框架断开DB链接，其他框架查看文档框架。 基于swoole框架的不需要
            echo "Worker#{$workerId} is started" . PHP_EOL;
            echo "共" . $cycleIndex . "页，开始第：" . $queue['index'] . '页' . PHP_EOL;

            //业务数据逻辑处理start
            $errorQuestion = [];
            $questionSql = $this->questionStl();
            $questionData = $questionSql->select(['question.question_id',
                'question.question_no', 'question_analysis', 'question_content'])
                ->forPage($queue['index'], $this->limit)
                ->groupBy('question_no')
                ->get();
            if ($questionData->isNotEmpty()) {
                $errorQuestion = $this->existsImage($questionData, $errorQuestion);
            }

            if (count($errorQuestion)) {
                //这是laravel5.5日志处理 ，其他版本框架请查看对应文档。
                $monolog = \Log::getMonolog();
                $monolog->popHandler();
                //写入错误题目到日志
                \Log::useFiles(storage_path('logs/error_invalid_path.log'));
                \Log::info("需要手动修复：");
                \Log::info(array_values($errorQuestion));
                $monolog = \Log::getMonolog();
                $monolog->popHandler();
                \Log::useFiles(storage_path('logs/error_path.log'));
            }

            //业务数据逻辑处理end

        });

        $pool->on("WorkerStop", function (Pool $pool, $workerId) use ($table, $cycleIndex) {
            $worker_status = $table->get($workerId, "worker_status");
            if ($worker_status == 1) {
                //放回worker进程 原子自增
                $table->incr("queue", "task_sum", $incrby = 1);
                $table->set($workerId, ['worker_status' => 2]);
            }
            $queue = $table->get("queue");
            if ($queue['index'] >= $cycleIndex) {
                //等待最后一个进程完成处理
                if ($queue['task_sum'] >= $this->workerNum) {
                    $pool->shutdown(); //关闭进程池，所有子进程结束运行。
                }
                return;
            }
            echo "Worker#{$workerId} is stopped" . PHP_EOL;
        });
        $pool->start();
    }

    /**
     * 数据查询
     * @return mixed
     */
    private function questionStl()
    {
        //题目查询
        return Question::query()
            ->where('question_status', 1)
            ->when($this->subject_id, function ($query) {
                $query->where('subject_id', $this->subject_id);
            })
            ->when($this->updated_at, function ($query) {
                $query->where('updated_at', '>=', $this->updated_at);
            });
    }

    /**
     * @param $questionData
     * @param $errorQuestion
     * @return mixed
     */
    private function existsImage($questionData, $errorQuestion)
    {
        //处理题目图片数据
        return $errorQuestion;
    }
}

```