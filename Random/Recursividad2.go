"1. Dado un arreglo de 10 elementos, desarrollar una función recursiva que encuentre
el valor máximo en el arreglo. La función debe comparar el elemento actual con el
máximo encontrado hasta el momento y continuar recursivamente con el resto del
arreglo. Puede definir la cantidad de parámetros que desee"

Accion recursiva2 (numeros: arreglo[1...10] de enteros) es
 Ambiente
	Funcion maximo(a: arreglo[1...10] de enteros; max,indice: entero): entero 
		si indice = 11 entonces  
			maximo:= max
		sino 
			si a[indice] > max entonces 
				maximo:= maximo(a,a[indice],indice+1)
			sino 
				maximo:= maximo(a,max,indice+1)
			Fsi
		Fsi
	FFuncion

	Proceso 
		Esc("El maximo valor del arreglo es: ",maximo(numeros,numeros[1],1))
	FProceso
FAccion