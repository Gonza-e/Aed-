Accion invertir (Arr1,Arr2: arreglo[1...10] de enteros) es 
 Ambiente

	Procedimiento invertir(A,B: arreglo[1...10] de enteros; indA, indB: entero) 
		si indA < 11 entonces 
			invertir(A,B,indA + 1,indB - 1)
		fsi

		B[indB]:= A[indA]
	FP 

	Proceso 

		invertir(Arr1,Arr2,1,10)

	FProceos 
FinAccion
