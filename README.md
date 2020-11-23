# yinyang 是一个公历和农历转换的服务

_农历和阴历并不是同一个东西，相差几天。为了更好起名字，这里用了 yin 和 yang。_

公历和农历并不是通过算法算出来的，而是天文台预测的。

香港天文台公共和农历对照表：https://www.hko.gov.hk/tc/gts/time/conversion1_text.htm

## 设计

1. 从香港天文台抓取源数据到 rawdata 目录；
2. 启动程序时，解析数据，建立映射关系，做成一个服务；

## API 列表

- `/api/v1/years` 年份列表
- `/api/v1/years/${year}/months` 某年的月份列表
- `/api/v1/years/${year}/months/${month}/${leap}/days` 某年某月的农历日历表
- `/conv/yang-yin/${year}/${month}/${day}` 公历转农历
- `/conv/yin-yang/${year}/${month}/${leap}/${day}` 农历转公历

## 项目状态

逻辑比较简单，已经生产化。
