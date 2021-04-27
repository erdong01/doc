## SOAR简单使用
#### 简介：
SOAR(SQL Optimizer And Rewriter) 是一个对 SQL 进行优化和改写的自动化工具。 由小米人工智能与云平台的数据库团队开发与维护。

#### 安装

```shell
composer require guanguans/soar-php --dev
```

```php

namespace App\Utils;

use Guanguans\SoarPHP\Soar;
use Illuminate\Support\Facades\DB;

class SoarUtil
{

    public static function score()
    {
        $connectionName = DB::getDefaultConnection();
        $databaseConfig = config('database.connections')[$connectionName];
        $dir = base_path();
        $config = [
            // 下载的 soar 的路径
            '-soar-path' => $dir . '/temp/soar',
            // 测试环境配置
            '-test-dsn' => [
                'host' => $databaseConfig['host'],
                'port' => $databaseConfig['port'],
                'dbname' => $databaseConfig['database'],
                'username' => $databaseConfig['username'],
                'password' => $databaseConfig['password'],
            ],
            // 日志输出文件
            '-log-output' => $dir . '/temp/soar.log',
            // 报告输出格式: 默认  markdown [markdown, html, json]
            '-report-type' => 'markdown',
        ];
        $soar = new Soar($config);
        $markdownStr = "";
        foreach (\DB::getQueryLog() as $valueD) {
            $sql = $valueD['query'];
            foreach ($valueD['bindings'] as $replace) {
                $value = is_numeric($replace) ? $replace : "'" . $replace . "'";
                $sql = preg_replace('/\?/', $value, $sql, 1);
            }
            $res = $soar->syntaxCheck($sql);
            if ($res) {
                $markdownStr .= $res;
            }
            $res = $soar->score($sql);
            $markdownStr .= $res;
            $res = $soar->mdExplain($sql);
            $markdownStr .= $res;
        }
        //检测报告路径
        $filepath = $dir . "/temp/soar.md";
        $fopen = fopen($filepath, "wb") or die("文件不存在");
        fwrite($fopen, $markdownStr);
        fclose($fopen);
        system('mdv ' . $filepath);
    }
}

```