def reverso(num):
    if num > 0:
        print(num % 10)
        reverso(num // 10)

reverso(254)