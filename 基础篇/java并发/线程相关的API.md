# java.lang.Object

## notify()
从在这个对象上wait着的线程中选择一个并唤醒它。

如果当前线程没有获取这个对象的锁（没有持有监视器）,在运行时时异常java.lang.IllegalMonitorStateException会被抛出。

## notifyAll()
唤醒在这个对象上wait着的全部线程

如果当前线程没有获取这个对象的锁(没有持有监视器)，在运行时时异常java.lang.IllegalMonitorStateException会被抛出。

## wait() throws IterruptedException
让当前线程（调用wait方法的线程）wait。不会发生超时。

与sleep方法不同，当前线程获取的锁会被释放。

如果当前线程没有获取这个对象的锁(没有持有监视器)，在运行时时异常java.lang.IllegalMonitorStateException会被抛出。

如果其他线程interrupt了当前线程，异常java.lang.InterruptedException会被抛出，中断状态也会被清除。

## wait(long mills) throws IterruptedException
让当前线程（调用wait方法的线程）wait。mills是发生超时的时间（以毫秒为单位）。如果是wait(0),则不会发生超时。

与sleep方法不同，当前线程获取的锁会被释放。如果mills为负数，在运行时异常java.lang.IllegalArgumentException会被抛出。

当正在wait时，如果其他线程interrupt了当前线程，异常java.lang.InterruptedException会被抛出，中断状态也会被清除。

## wait(long mills， int nanos) throws IterruptedException
让当前线程（调用wait方法的线程）wait。1000000*mills + nanos 是发生超时的时间（以纳秒为单位）。如果是wait(0，0),则不会发生超时。

与sleep方法不同，当前线程获取的锁会被释放。如果mills为负数或者nanos值不在0至999999之间，在运行时异常java.lang.IllegalArgumentException会被抛出。

如果当前线程没有获取这个对象的锁(没有持有监视器)，在运行时时异常java.lang.IllegalMonitorStateException会被抛出。

如果其他线程interrupt了当前线程，异常java.lang.InterruptedException会被抛出，中断状态也会被清除。

# java.lang.Runnable接口

## run()
在创建java.lang.Thread的实例时，如果指定实现了Runnable接口的类的实例，那么线程启动后会调用run方法。

# java.lang.Thread(implements Runnable)

## Thread()
创建java.lang.Thread的实例

## Thread(Runnable target)
指定一个实现了java.lang.Runnable接口的类的实例，创建java.lang.Thread的实例

## Thread(Runnable target, String name)
指定一个实现了java.lang.Runnable接口的类的实例和名字，创建java.lang.Thread的实例

## Thread(String name)
指定一个名字，创建java.lang.Thread的实例

## Thread(ThreadGroup group, Runnable target)
指定一个线程组，创建java.lang.Thread的实例

## Thread(ThreadGroup group, Runnable target, String name)
指定一个线程组，一个实现了java.lang.Runnable接口的类的实例，以及一个名字，创建java.lang.Thread的实例

## Thread(ThreadGroup group, String name)
指定一个线程组和一个名字，创建java.lang.Thread的实例

## MIN_PRIORITY
表示可以设置的线程的最低优先级

## NORM_PRIORITY
表示可以设置的默认优先级

## MAX_PRIORITY
表示可以设置的最高优先级

## holdsLock(Object obj)
如果当前线程持有obj的监视器，则本方法返回true

## interrupt()
中断本线程（this)
请注意，这里是中断与this对应的线程，并不一定会中断与Thread.CurrentThread()对应的线程

## boolean interrupted()
判断当前线程（调用interrupted方法的线程）是否处于中断状态。调用本方法后，当前线程将不再处于中断状态

## boolean isInterrupted()
判断本线程（this)是否处于中断状态。即使调用本方法，本线程（this)的中断状态也不会变化。
