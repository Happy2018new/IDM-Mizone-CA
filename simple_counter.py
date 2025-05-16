c1 = []  # 1代言人影响 (黄)
c2 = []  # 2购买意愿 (蓝)
c3 = []  # 3产品功能关联 (绿)
c4 = []  # 4品牌活动 (胜蓝)
c5 = []  # 5产品体验 (灰)
c6 = []  # 6服务渠道 (紫)

while True:
    token = input()
    if token == "q":
        break
    try:
        t = int(token[0])
    except Exception:
        continue
    if t == 1:
        c1.append(token[1:])
    if t == 2:
        c2.append(token[1:])
    if t == 3:
        c3.append(token[1:])
    if t == 4:
        c4.append(token[1:])
    if t == 5:
        c5.append(token[1:])
    if t == 6:
        c6.append(token[1:])

print(set(c1))
print(set(c2))
print(set(c3))
print(set(c4))
print(set(c5))
print(set(c6))
