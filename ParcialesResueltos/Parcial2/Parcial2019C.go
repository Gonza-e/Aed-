Accion Ejercicio1 es 
 Ambiente 
	Fecha = registro 
		dia: N(2)
		mes: N(2)
		anio: N(4)
	FinRegistro
	
	MILLAS = registro 
		dni: N(8)
		millas: N(10)
		ult_carga: Fecha
	FinRegistro
	mill: archivo de MILLAS indexado por dni 
	reg_m: MILLAS 

	DESTINOS = registro 
		origen: N(1)
		destino: N(1)
		millas: N(10)
		duracion: N(2)
	FinRegistro
	dest: archivo de DESTINOS indexado por origen y destino 
	reg_d: DESTINOS
	DNI, ORIGEN, DESTINO, millas1, c_nuevos, cod, cod2: entero

	Procedimiento descontar_millas() es 
		reg_m.millas:= reg_m.millas - reg_d.millas
		Regrabar(mill,reg_m)
	FinProcedimiento

	Procedimiento suma_de_millas() es 
		Esc("Ingrese la cantidad de millas que desea cargar")
		Leer(millas1)
		reg_m.millas:= reg_m.millas + millas1
		Regrabar(mill,reg_m)
	FinProcedimiento

	Procedimiento sumar_millas() 
		Esc("Ingrese codigo de operacion (1: sumar millas, 2: salir)")
		Leer(cod2)
	 	Si cod2 = 1 entonces 
		 	suma_de_millas() 
		FinSi
	FinProcedimiento

	Proceso 
		Abrir E/S (mill); Abrir E/(dest)
		DNI:=0, ORIGEN:=0, DESTINO:=0; c_nuevos:=0, cod:=0; cod2:=0

		Repetir 

			Esc("Ingrese origen y destino")
			Leer(ORIGEN); Leer(DESTINO)
			reg_d.origen:= ORIGEN
			reg_d.destino:= DESTINO
			Leer(dest,reg_d) 

		 	Si EXISTE entonces 
				Esc("Ingrese dni del pasajero")
				Leer(DNI)
				reg_m.dni:= DNI 
				Leer(mill,reg_m)
			 	Si EXISTE entonces
					Esc("Ingrese codigo de operacion (1: calcular millas, 2: salir)")
					Leer(cod)
				 	Si cod = 1 entonces 
						Segun reg_m.millas hacer 
							> reg_d.millas: descontar_millas()
							< reg_d.millas: sumar_millas()
						FinSegun 
					FinSi
				Sino 
				 	c_nuevos:= c_nuevos+1
				FinSi
			Sino 
			 	Esc("El destino u origen son incorrectos")
			FinSi
		Hasta que (cod=2) 

		Esc("La cantidad de pasajeros nuevos es de",c_nuevos)
		Cerrar(mill)
		Cerrar(dest)
	FinProceso
FinAccion

Accion Ejercicio2 (A: arreglo[1...10] de AN(3)) es 
	Ambiente

	Fecha = registro
		dia: N(2)
		mes: N(2)
		anio: N(4)
	FinRegistro

	VIAJES = registro
		fecha: Fecha 
		origen: N(2)
		destino: N(2)
		m_compradas: N(10)
	FinRegistro
	via: archivo de VIAJES 
	reg_v: VIAJES
	B: arreglo[1...5,1...11] de enteros 
	menor, mayor, i, j, d_may, d_men, t_may: entero 

	Proceso 
		Abrir E/(via)
		Leer(via,reg_v)
		menor:= HV, mayor:=0 i:=0; j:=0; t_may:=0; d_men:=0; d_may:=0

		Para i:= 1 a 5 hacer
			Para j:= 1 a 11 hacer  
				B[i,j]:= 0
			FinPara
		FinPara

	 	Mientras NFDA(via) hacer 
			j:= reg_v.destino 
			Segun reg_v.fecha.mes hacer 
				1...3: i:= 1 
				4...6: i:= 2 
				7...9: i:= 3 
				10...12: i:= 4 
			FinSegun
			B[i,j]:= B[i,j] + reg_v.m_compradas
			B[5,j]:= B[5,j] + reg_v.m_compradas
			B[i,11]:= B[i,11] + reg_v.m_compradas
			Leer(via,reg_v)
		FinMientras

		Para i:= 1 a 5 hacer 
			Para j:= 1 a 11 hacer 
		 	 	Si B[i,j] > mayor entonces
					mayor:= A[i,j]
					t_may:= i 
					d_may:= j 
				FinSi
				Si B[i,j] < menor entonces 
					menor:= B[i,j]
					d_men:= j
				FinSi
			FinPara 
		FinPara 

		Para i:= 1 a 10 hacer 
			Esc("El destino con mas millas compradas es",A[d_may])
			Esc("El destino con menos millas compradas es",A[d_men])
	    FinPara 

		Esc("El trimestre con mas millas compradas es",t_may)
		Cerrar(via)
	FinProceso
FinAccion




	 	 

	 


						 
					  