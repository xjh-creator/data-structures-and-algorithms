##解决的问题：为什么栈作为一种“操作受限”的数据结构，存在的意义？
栈很多操作都可以由数组或链表来说实现，但它们暴露出太多的多余的操作接口，容易导致一些问题。不同的数据结构有不同的适应场景。

##特性
1. 先进后出

##应用场景：
###一、浏览器中上下页的切换
定义两个A跟B栈，A栈负责把浏览的页面一一推入存储，最上面的页面就是当前访问的页面，当需要翻到上一页时，就把最上面的页面丢到B栈
中，而A栈最上面的页面就是翻到上一页的页面中。

###二、函数的调用
访问函数时，会把当中的变量或表达式等一一压入栈中，处理完就会弹出来。

###三、计算表达式的应用
访问一个表达式如：1+3*5-4 时，定义俩个栈A跟B,A栈负责存数字，B栈负责存操作符，每次压入操作符前，先把操作符跟B栈中最上层的操作
符比较，如果大于栈中的操作时，不压入操作符，而是从A栈中弹出两个数字跟准备压入栈的操作符进行计算，把计算结果压入A栈，然后继续扫描
表达式。

###四、（）的匹配