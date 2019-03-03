### a-star-redblob
This is a Golang implementation of RedBlobGames A* algorithm from 
[tutorial article](https://www.redblobgames.com/pathfinding/a-star/implementation.html).

To run, execute:
```bash
$ go build . 
$ ./a-star-redblob
```

Examplary terminal output: 
```
 .  .  .  .  .  .  .  .  .  .   | S start                                               
 .  .  .  .  .  .  .  .  .  .   | G goal                                                
 .  .  .  .  ^  ^  ^  .  .  .   | * path step                   
 .  .  .  .  ^  ^  ^  .  .  .   | ^ woods (costs 5 to step into)
 .  G  .  .  ^  ^  ^  .  .  .   | . roads (costs 1 to step into)
 .  *  .  .  ^  ^  ^  .  .  .   | X walls (cannot step into)    
 *  *  .  .  ^  ^  ^  .  .  .    
 *  X  X  X  ^  ^  ^  .  .  .    
 *  X  X  X  *  *  *  S  .  . 
 *  *  *  *  *  .  .  .  .  . 
```

Original examples were written in Python/C++/C#. The logic is almost exactly the same as in the article,
I only translated it to Go. I did it while learning A* for my AI course at the university.

Anyone studying RedBlobGames tutorial can look here for reference.
I highly recommend getting familiar with their work.

*License MIT*

