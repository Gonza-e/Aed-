Accion invertir (Arr1: arreglo[1...10] de enteros) es 
 Ambiente

	Procedimiento invertir(A: arreglo[1...10] de enteros; indA,x: entero) 
		si indA < 11 entonces 
			x:= A[11-indA]
			invertir(A,indA+1,x)
			A[indA]:= x
		fsi
	FP 

	Proceso 

		invertir(Arr1,Arr2,1,10)

	FProceso 
FinAccion
