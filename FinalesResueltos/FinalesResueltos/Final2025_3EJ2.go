Accion final2025 es 
 Ambiente 
	nodo = registro 
		valor: entero 
		prox: puntero a nodo 
	Freg
	p,s,q,prim,ant: puntero a nodo 
	
	funcion numeroPrometido(num1,num2,num3,suma: entero): booleano 
		si num3 = 0 entonces 
			si (num2 + 1) = suma entonces 
				numeroPrometido:= verdadero 
			sino 
				numeroPrometido:= falso 
			fsi 
		sino 
			si (num1 mod num3) = 0 entonces 
				numeroPrometido:= numeroPrometido(num1, num2, num3 - 1, suma + num1)
			sino 
				numeroPrometido:= numeroPrometido(num1, num2, num3 - 1, suma)
			fsi 
		fsi 
	FFuncion

	numero: entero 

	Proceso
		prim:= nil 
		repetir 
			Esc("Ingrese un valor")
			Leer(numero)

			si numero > 0 entonces
				Nuevo(q); *q.valor:= numero 
				*q.prox:= prim 
				prim:= q 
			fsi 
		hasta que (numero < 0)

		p:= prim, s:= *prim.prox 

		Mientras p <> nil hacer 
			Mientras s <> nil hacer
				si numeroPrometido(*p.valor,*s.valor,*p.valor - 1,0) y numeroPrometido(*s.valor,*p.valor,*s.valor - 1,0) entonces 
					Esc("Los numeros son prometidos")
				sino 
					Esc("Los numeros no son prometidos")
				fsi 
				s:= *s.prox 
			FM 
			p:= *p.prox 
			s:= p; s:= *p.prox
		FM 
	FProceso 
FAccion 