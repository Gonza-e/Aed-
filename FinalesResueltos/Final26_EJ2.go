Accion ej2 es 
	Ambiente 

	cuil, clave, k, res: entero 
	A: arreglo[1...10] de entero 
	bandera: booleano

	USUARIO = registro 
		cuil: N(11)
		clave: N(10)
	Freg 
	arch: archivo de USUARIO indexado por cuil y clave 
	reg: USUARIO 

	Procedimiento cargarClave()
		Para i:= 1 hasta 10 hacer 
			A[11-i]:= clave mod 10 
			clave:= clave div 10 
		FPara 
	FProcedimiento

	Procedimiento verificarClave(bandera) 
		Para i:= 1 hasta 10 hacer 
			si i mod 2 <> 0 entonces 
				A[i]:= A[i]+A[i+1]
			Fsi 

			res:= A[i] div k 
			si res mod 2 <> 0 entonces 
				A[i]:= res 
			Fsi 
			k:= k+1 
		FPara 

		clave:= 0 
		Para i:= 1 hasta 10 hacer 
			clave:= clave*10 + A[i] 
		FPara 

		reg.cuil:= cuil; reg.clave:= clave; Leer(arch,reg)
		si EXISTE entonces 
			bandera:= verdadero
		sino 
			bandera:= falso 
		Fsi 
	FProcedimiento

	Proceso 
		Abrir E/(arch)
		k:= 3
		Esc("Ingrese su cuil:"); Leer(cuil)
		Esc("Ingrese clave:"); Leer(clave)

		cargarClave()
		verificarClave(bandera)

		si bandera entonces 
			Esc("El usuario existe")
		sino 
			Esc("El usuario no existe")
		Fsi 

		Cerrar(arch)
	FProceso 

FAccion
			

