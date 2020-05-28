### API 支持多版本运行

到移动互联网万物互联时代，APP还是蓬勃发展年代。原生APP+H5各种版本。
在产品经理+老板各种要求，各种需求驱动下。
如果有个不成熟的产品经理或者hould 不住老板的情况下各种需求一顿猛如虎的操作版本不断叠加，那就糟糕了，代码是会越来越乱了。

生为后端人员，为了保住代码质量和生活品质。这是需要面向接口编程来助阵了。

### 环境：
php 7.4
laravel 5.5
swoole 4.5+



#### 1.获取uri版本号。
```php
if (!function_exists('getVersion')) {
    /**
     * 获取请求版本
     * @return mixed|string
     */
    function getVersion()
    {
        $version = "v1";
        $urlSegments = \Request::segments();
        if (isset($urlSegments[0])) {
            $version = $urlSegments[0];
        }
        return $version;
    }
}
```

#### 2.面向接口需要用到laravel  ServiceProvider(服务提供器)

https://learnku.com/docs/laravel/5.5/providers/1290

```php
namespace App\Providers\Services\Common;

use App\Services\Contract\Common\IExamService;
use Illuminate\Support\ServiceProvider;

class ExamServiceProvider extends ServiceProvider
{
    /**
     * 是否延时加载提供器。
     *
     * @var bool
     */
    protected $defer = true;

    /**
     * Bootstrap the application services.
     *
     * @return void
     */
    public function boot()
    {
        //
    }

    /**
     * Register the application services.
     *
     * @return void
     */
    public function register()
    {
        //绑定服务
        $this->app->bind(IExamService::class, function () {
            return app('App\Services\Common\\' . getVersion() . '\ExamService');
        });
    }

}
```