# 正则表达式

[Javascript正则表达式](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Guide/Regular_Expressions)

## 简单模式

无特殊字符的字符串匹配

```javascript
var regexp = new RegExp("abc");
regexp.exec("aaaaaaabcccccc");
```

## 特殊字符

<table>
  <tr></tr>
    <td>字符</td>
    <td>含义</td>
  <tr><td>\</td><td>
    将其后的字符转义为字面量。<br/>
    例如模式/a*/可以匹配0个或多个a，而模式/a\*/将*的特殊性移除，可以匹配"a*"这样的字符串
  </td></tr>
  <tr><td>^</td><td>
    匹配输入的开始
  </td></tr>
  <tr><td>$</td><td>
    匹配输入的结束
  </td></tr>
  <tr><td>*</td><td>
    匹配前一个表达式0次或者多次，等价于{0,}
  </td></tr>
  <tr><td>+</td><td>
    匹配前一个表达式1次或者多次，等价于{1,}
  </td></tr>
  <tr><td>{n}</td><td>
    n是一个正整数，匹配前一个字符恰好发生了n次
  </td></tr>
  <tr><td>{n,m}</td><td>
    n和m都是整数，匹配前一个字符至少发生n次，至多m次
  </td></tr>
  <tr><td>?</td><td>
    匹配前一个表达式0次或者1次，等价于{0,1}<br/>
    如果紧跟在任何量词* + ? {}后面，将会使量词变为非贪婪的（匹配尽量少的字符）
  </td></tr>
  <tr><td>.</td><td>
    匹配除换行符之外的任意其他字符
  </td></tr>
  <tr><td>(x)</td><td>
    匹配x并且记住匹配项，括号()被称为捕获括号。<br/>
    模式/(foo) (rar) \1 \2/，其中的\1就表示第1个被捕获并且记住的匹配单词"foo"，\2表示第2个被捕获并且记住的匹配单词"rar"。\1 \2 .. \2这种写法使用在正则表达式的匹配环节。$1 $2 .. $n这种写法是用在正则表达式的替换(replace)环节。例如："mike jordan".replace(/(\w+)\s(\w+)/, "$2 $1")。
  </td></tr>
  <tr><td>(?:x)</td><td>
    匹配x但是不记住匹配项，这种叫做非捕获括号。<br/>
    "foofoofoo".
    /(?:foo)+/，这个可以匹配foofoofoo
  </td></tr>
  <tr><td>x(?=y)</td><td>
    先行断言，匹配x仅仅当x后面紧跟y。<br/>
    /Jack(?=Sprat|Frost)/匹配‘Jack’仅仅当它后面跟着'Sprat'或者是‘Frost’。但是‘Sprat’和‘Frost’都不是匹配结果的一部分。
  </td></tr>
  <tr><td>(?&lt=y)x</td><td>
    后行断言，匹配x仅仅当x前面是y。
  </td></tr>
  <tr><td>x(?!y)</td><td>
    正向否定查找，匹配x仅仅当后面跟的不是y。
  </td></tr>
  <tr><td>x|y</td><td>
    匹配x或者y。
  </td></tr>
  <tr><td>[xyz]</td><td>
    字符集合，匹配方括号中的任意字符，(. *)这样的特殊字符在方括号中不需要转义。我们可以使用破折号(-)来表示范围。例如，[a-z]。
  </td></tr>
  <tr><td>[^xyz]</td><td>
    反向字符集，不匹配方括号中的任意字符。
  </td></tr>
  <tr><td>\b</td><td>
    匹配一个词的边界，表示词的开头或者是结尾。
  </td></tr>
  <tr><td>\B</td><td>
    TODO
  </td></tr>
  <tr><td>\cX</td><td>
    如果x是A-Z之间的字符，匹配字符串中的一个控制符。<br/>
    /\cM/匹配字符串中的control-M
  </td></tr>
  <tr><td>\d</td><td>
    匹配一个数字，等价于[0-9]
  </td></tr>
  <tr><td>\D</td><td>
    匹配一个非数字，等价于[^0-9]。
  </td></tr>
  <tr><td>\f</td><td>
    匹配一个换页符。
  </td></tr>
  <tr><td>\n</td><td>
    匹配一个换行符。
  </td></tr>
  <tr><td>\r</td><td>
    匹配一个回车符。
  </td></tr>
  <tr><td>\s</td><td>
    匹配一个空白符，包括空格，制表位，换页符，换行符。
  </td></tr>
  <tr><td>\S</td><td>
    匹配一个非空白字符。
  </td></tr>
  <tr><td>\w</td><td>
    匹配一个单字字符（字母，数字，下划线）,等价于[a-zA-Z0-9_]。
  </td></tr>
  <tr><td>\W</td><td>
    匹配一个非单字字符，等价于[^a-zA-Z0-9_]。
  </td></tr>
</table>
