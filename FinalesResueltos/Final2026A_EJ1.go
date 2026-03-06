Accion ej1 es 
	Ambiente 

	sec: secuencia de caracteres 
	s: caracter 

	nodop = registro 
		prin: puntero a nodo 
		cr: N(1)
		id: N(6)
		prox: puntero a nodop 
	Freg
	prim,p,q,ant: puntero a nodop 

	nodo = registro 
		n: N(3)
		d: AN(1)
		prox: puntero a nodo 
	Freg 
	prim1,p1,q1,ult1: puntero a nodo 

	vocales: {"a","e","i","o","u","A","E","I","O","U"}
	digitos: {"0","1","2","3","4","5","6","7","8","9"}
	cons: {"b","c","d","f","g","h","j","k","l","m","n","p","q","r","s","t","v","w","x","y","z"}
	cri,ide,cont,c1,c2,voc: entero 
	let,sig: AN 

	Funcion valor(x:AN): entero 
		segun x hacer 
			"a": valor:= 1 
			"e": valor:= 2 
			"i": valor:= 3 
			"o": valor:= 4 
			"u": valor:= 5 
			"A": valor:= 1 
			"E": valor:= 2 
			"I": valor:= 3 
			"O": valor:= 4 
			"U": valor:= 5
		Fsegun
	FF
	
	Funcion letra(x:entero): AN 
		segun x hacer 
			1: letra:= "a"
			2: letra:= "e"
			3: letra:= "i"
			4: letra:= "o"
			5: letra:= "u"
		Fsegun 
	FF 

	Procedimiento obtInfo()
		Para i:= 1 hasta 6 hacer 
			ide:= ide*10 + conv(s)
			Avz(sec,s)
		FPara 
		Para i:= 1 hasta 7 hacer 
			Avz(sec,s)
		FPara
		cri:= conv(s)
		Mientras s <> "-" o s <> "+" hacer 
			guardar()
			Avz(sec)
		FM 
		sig:= s 
		Avz(sec,s)
	FP 
	
	Procedimiento verificar()
		p1:= prim1 
		Mientras p1 <> nil hacer 
			si sig = "-" entonces 
				si *p1.n en digitos entonces 
					si conv(*p1.n) mod 2 = 0 entonces 
						*p1.n:= conv(*p1.n)*2 
					sino 
						*p1.n:= conv(*p1.n)**3 
					Fsi 
					c2:= c2+1
				sino 
					si s no en cons y s no en vocales entonces 
						*p1.d:= "!"
						c1:= c1+1
					Fsi 
				Fsi 
			sino 
				si *p1.d en cons o *p1.d en vocales entonces 
					si *p1.d en vocales entonces 
						reemplazar()
					sino 
						*p1.d:= ASCII(*p1.d)
					Fsi 
					c2:= c2+1
				sino 
					si ((*p1.d no en cons) y (*p1.d no en vocales)) y (*p1.n = nil) hacer 
						*p1.d:= "!"
						c1:= c1+1
					Fsi 
				Fsi 
			Fsi 
			p1:= *p1.prox 
		FM 
	FM 

	Procedimiento guardar()
		Nuevo(q1)
		si s en digitos entonces 
			*q1.n:= conv(s)
		sino 
			*q1.d:= s 
		Fsi 

		si prim1 = nil entonces 
			*q1.prox:= prim1 
			prim1:= q1 
			ult1:= q1 
		sino
			*q1.prox:= *ult1.prox 
			*ult1.prox:= q1 
			ult1:= q1 
		Fsi 
	FP 

	Procedimiento reemplazar()
		voc:= valor(*p1.d)
		voc:= (voc-2) mod 5 
		let:= letra(voc)
	FP

	Procedimiento cargaOrdenada()
		Nuevo(q); *q.id:= ide; *q.cr:= cri; *q.prin:= prim1; prim1:= nil 

		si prim = nil entonces 
			*q.prox:= prim 
			prim:= q 
		sino 
			ant:= nil; p:= prim 
			Mientras p<>nil y ((*q.cri<*p.cri) o ((*q.cri=*p.cri) y (*q.id<*p.id))) hacer
				ant:= p 
				p:= *p.prox 
			FM 
			
			si prim = p entonces 
				*q.prox:= prim 
				prim:= q 
			sino 
				*ant.prox:= q 
				*q.prox:= p 
			Fsi 
		Fsi 
	FP

	Proceso 
		Arr(sec); Avz(sec,s)
		cri:= 0; ide:= 0; cont:= 0; c1:= 0; c2:= 0; voc:= 0
		sig:= ""; let:= ""

		Mientras NFDS(sec) hacer 
			obtInfo()
			verificar()
			cargaOrdenada()
			ide:= 0
		FM 
		Cerrar(sec)

		p:= prim
		Mientras p<>nil hacer 
			p1:= prim1
			Mientras p1<>nil hacer 
				si *p1.d <> nil hacer 
					Esc(*p1.d)
				sino 
					Esc(*p1.n)
				Fsi 
				p1:= *p1.prox 
			FM 
			p:= *p.prox 
		FM 
		cont:= c1+c2 
		porc:= (c1/cont)*100; Esc("Porcentaje de caracteres no validos: ",porc,"%")

		si (c1+c1*0.10) > c2 entonces 
			Esc("Sistema comprometido")
		sino 
			si c1 = c2 entonces 
				Esc("Sistema en alerta")
			sino 
				si c1 < (c2 div 2) entonces 
					Esc("Sistema seguro")
				Fsi 
			Fsi 
		Fsi 
	FProceso 
FAccion



