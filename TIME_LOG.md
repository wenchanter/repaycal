|    Area     | Expect |  Actual   | Reason                                                                                                                                                   |
|:-----------:|:------:|:---------:|----------------------------------------------------------------------------------------------------------------------------------------------------------|
|  proto/API  | 30 min |  20 min   |                                                                                                                                                          |
| calculation | 30 min |  15 min   | AI helped me finish the method and then I reviewed it. It worked well.                                                                                   |
| persistence | 30 min |  40 min   | Understood the framework’s model feature, but was stuck on the framework’s connection driver                                                             |
|   docker    | 30 min |  1 hour   | Used the go-zero’s scaffold to create the Dockerfile, and then moved it to the root dir. It didn’t work well.                                            |
|     k8s     | 30 min |  1 hour   | Some config problem cost too much time.                                                                                                                  |
|   tests     | 1 hour | 2.5 hours | UT: Understood the framework when debugging. It took 40 min to connect the model, logic and service modules.                                             |
|  | |           | Local environment integration testing: 20 min                                                                                                            |
|  |  |           | docker_run.sh file testing and docker environment integration testing: 30 min                                                                            |
|  |  |           | k8s_run.sh file testing and k8s environment integration testing: 50 min. The image cache couldn’t be deleted in k8s until I clean the whole k8s cluster. |
| docs | 30 min |  1 hour   | As clear as possiable                                                                                                                                    |


Extra time: 3 hours  
1. Learned the overall picture of the loan business, identified which domain this feature belongs to, and thought about how to extend the calculation. 30 min  

2. Considered the project need to support many production features, and basically it should contain service registration and service discovery, load balance, metrics, tracing. I decided to use a framework. At ByteDance, we used CloudWeGo, but it was quite heavy and complex. It uses Thrift and Consul by default, so perhaps I have had to solve many problems if I wanted to use protobuf with it.  
Finally I decided to use go-zero, which is popular in China. However, it was the first time I use it, so I spent more than 2 hours to learning the framework and understanding how it works.


Need to improve
1. I need more time to do more testing, which includes increasing code coverage and thinking about the edge cases. It would be better if I could introduce an automated testing framework to check code coverage and do the automated test.
2. I need to understand go-zero more deeply, and read more core source code to see how it works, so that if there was an incident I can diagnose it quickly.
3. I need to learn more about the loan business in order to design a more extensive system.   

How I use AI
1. I used AI to learn about the loan business.
2. I used AI to help me create the Dockerfile and k8s deployment files, and asked AI when I dealt with deployment problems.
3. I used AI to generate some specific methods and most of the test files, then reviewed the code.
4. I ask AI to correct grammatical mistakes where I wrote the doc.
