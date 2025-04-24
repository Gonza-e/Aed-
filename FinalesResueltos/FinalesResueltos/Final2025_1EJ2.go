Accion Ejercicio2 es 
 Ambiente 
	sec: secuencia de caracteres 
	v: caracter 
	
	CAMIONES = registro 
		nro_patente: AN(6)
		chofer: AN(30)
	FReg  

	nodo = registro 
		datos: CAMIONES
		cant_entregas: N(7)
		prox: puntero a nodo 
	FReg 

	p,prim,q,ant: puntero a nodo 

	arch: archivo de CAMIONES 
	reg: CAMIONES	
	mayor, numero, i: entero 

	funcion capicua(num1, num2, num1_og: entero): booleano 
		si num1 = 0 entonces 
			si num1_og = num2 entonces 
				capicua:= verdadero 
			sino 
				capicua:= falso 
			fsi 
		sino 
			capicua:= capicua(num1 div 10, (num2*10) + (num1 mod 10), num1_og)
		fsi 
	FFuncion

	procedimiento cargaOrdenada() es 
		si prim = nil entonces 
			*q.prox:= nil 
			prim:= q 
		sino 
			p:= prim; ant:= nil 
			Mientras (p<>nil) y (*q.cant_entregas > *p.cant_entregas) hacer 
				ant:= p 
				p:= *p.prox 
			FM 

			si p = prim entonces 
				*q.prox:= prim  
				prim:= q 
			sino 
				*q.prox:= *ant.prox 
				*ant.prox:= q 
			fsi 
		fsi 
	FP

 Proceso 
		Abrir E/(arch); Leer(arch,reg)
		Arr(sec); Avz(sec,v) 
		mayor:= 0 

		Para i:= 1 a 10 hacer 
			top[i].cant_entregas:= 0 
		FPara 
		
		Mientras NFDS(sec) y NFDA(arch) hacer 
			Mientras v <> "*" hacer 
				Mientras v <> "+" hacer 
					numero:= numero + (conv_entero(v)*(10**i))
					i:= i + 1 
					Avz(sec,v)
				FM 
				Avz(sec,v) 

				si capicua(numero,0,numero) entonces 
					Nuevo(q); *q.datos:= reg; *q.cant_entregas:= numero 
					cargaOrdenada()
				fsi
				Leer(arch,reg); numero:= 0; i:= 0
			FM 
		FM 

		p:= prim 
		Para i:= 1 a 10 hacer 
			Esc("Top ",i,": ",*q.datos.nro_patente,"/",*q.datos.chofer)
			p:= *p.prox 
		FPara
		
		Cerrar(arch)
		Cerrar(sec)
	FProceso 
FAccion



	