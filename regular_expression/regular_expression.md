# 正则表达式

[Javascript正则表达式](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Guide/Regular_Expressions)

## 简单模式

无特殊字符的字符串匹配

```javascript
var regexp = new RegExp("abc");
var str = "aaaaaaabcccccc";
regexp.exec(str);

regexp = /aaa/;
str.match(regexp);
```

## [Special Character(特殊字符)](special_character.md)

## 正则表达式修饰符

- i ignore(不区分大小写)

```
var regex = /abc/i
['abc', 'Abc', "ABC"].forEach(str => { str.match(regex) }) // => true
```

- g global(全局匹配)

```
if there is no 'g', just return the first match element
if there is a 'g', it will return all matches
```

- m more(多行匹配)

```
if the string has line break \n and the regular expression has start or end symbol, g and m can make global match.
if only g, only match the first line. The m can match more lines.

var str = "abcggab\nabcoab";
var preg1 = /^abc/gm;  str.match(preg1)  // => ["abc", "abc"]
var preg2 = /ab$/gm;   str.match(preg2)  // => ["ab", "ab"]
```
