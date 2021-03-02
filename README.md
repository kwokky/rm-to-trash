# rm-to-trash

Go练手小项目，将文件移动到回收站。可代替`rm`命令。

## 构建
需要 `go>=1.11`，支持`go mod`
```bash
git clone https://github.com/kwokky/rm-to-trash
cd rm-to-trash
go build -o rm-to-trash
```

## 参数
- `-r  --recursive` 删除目录
- `-f  --force` 强制删除文件，不会提示

## 使用
默认移动到`$HOME/.Trash`目录下，兼容`rm`的`-r`、`-f`参数。

如果愿意的话也可以给把`rm`的alias设置为`rm-to-trash`。

```bash
# 删除当前目录下的file文件
rm-to-trash file

# 删除当前目录下的file1、file2文件
rm-to-trash file1 file2

# 强制删除当前目录下的file1、file2文件
rm-to-trash -f file1 file2

# 删除当前目录下的dir目录
rm-to-trash -r dir

# 强制删除当前目录下的dir1、dir2目录
rm-to-trash -rf dir1 dir2
```