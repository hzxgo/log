# log
the log plug-in is base on [logrus](github.com/sirupsen/logrus)

### Usage

* basic usage

    ```
    import (
        "github.com/hzxgo/log"
    )

    func main() {
        var username string
        username = "HeZhiXiong"

        log.Debugf("this is debug log | %s", username)
        log.Infof("this is info log | %s", username)
        log.Warnf("this is warn log | %s", username)
        log.Errorf("this is error log | %s", username)
        log.Fatalf("this is fatal log | %s", username)
        log.Panicf("this is panic log | %s", username)
    }
    ```

* init log by self
    ```
    import (
        "github.com/hzxgo/log"
    )

    func init() {
        logPath := "/data/logs/your_app_name/default.lgo"
        log.Init(true, 3, logPath) // it will write log to 'logPath' and delete three days ago logs
    }

    func main() {
        var username string
        username = "HeZhiXiong"

        log.Debugf("this is debug log | %s", username)
        log.Infof("this is info log | %s", username)
        log.Warnf("this is warn log | %s", username)
        log.Errorf("this is error log | %s", username)
        log.Fatalf("this is fatal log | %s", username)
        log.Panicf("this is panic log | %s", username)
    }
    ```

* set log level
you also can set log level by yourself like this ```log.SetLevel(log.INFO_LEVEL)```