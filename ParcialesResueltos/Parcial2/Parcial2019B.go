Accion Ejercicio1 (A: arreglo[1...200] de entero) es 
 Ambiente 
 Fecha = registro 
	 dia: N(2)
	 mes: N(2)
	 anio: N(4)
	FinRegistro
 TARJETA = registro 
	 dni: N(8)
	 n_cuenta: N(8)
	 credito: N(10)
	 ult_carga: Fecha
	FinRegistro
 tarje: archivo de TARJETA indexado por dni 
 reg_tar: TARJETA
 TURISTA = registro 
	 dni: N(8)
	 nombre: AN(100)
	 f_nacimiento: Fecha 
	 telefono: AN(100)
	 valido: logico 
	FinRegistro
 turis: archivo de TURISTA indexado por dni 
 reg_tur: TURISTA
 FACTURAS = registro
	 nro: N(10)
	 dni: N(8)
	 id_servicio: N(10)
	 monto: N(3)
	 f_carga: Fecha 
	FinRegistro
 factu: archivo de FACTURAS ordenado por nro y dni 
 reg_f: FACTURAS
 m_total, serv: entero 
 servicio: logico 
 Procedimiento cargar_monto() es 
	 reg_tar.credito:= reg_tar.credito + m_total
	 Regrabar(tarje,reg_tar)
	 Esc("El monto total es de",m_total)
	FinProcedimiento
 Procedimiento dar_baja() es 
	 reg_tar.ult_carga.dia:= 01
	 reg_tar.ult_carga.mes:= 01
	 reg_tar.ult_carga.anio:= 1999
	 Regrabar(tarje,reg_tar)
	 reg_tur:= falso 
	 Regrabar(turis,reg_tur)
    FinProcedimiento
 Proceso 
	 Abrir E/S (tarje); Abrir E/S (turis); Abrir E/(factu)
	 Leer(factu,reg_f)
	 m_total:=0;
	 Mientras NFDA(factu) hacer 
		 servicio:= falso
		 serv:= reg_f.id_servicio 
		 Para i:= 1 a 200 hacer  
		 	 	Si A[serv] = 1 entonces
				 servicio:= verdadero 
				FinSi 
			FinPara
			Si (reg_f.monto < 200) Y (servicio) entonces 
			 m_total:= ((reg_f.monto) DIV dolar2019)*0.50
			 reg_tar.dni:= reg_f.dni 
			 Leer(tarje,reg_tar)
			 	Si EXISTE entonces 
				 reg_tur.dni:= reg_f.dni 
				 Leer(turis,reg_tur)
				 	Si EXISTE entonces 
					 	Si m_total < 200000 entonces 
						 cargar_monto()
						Sino 
						 dar_baja()
						FinSi 
					FinSi
				FinSi
			FinSi
		 Leer(factu,reg_f)
		FinMientras
	 Cerrar(factu)
	 Cerrar(tarje)
	 Cerrar(turis)
	FinProceso
FinAccion

Accion Ejercicio2 (A: arreglo[1...23] de AN) es
 Ambiente 
	VIAJES = registro 
		dni: N(8)
		fecha: Fecha
		p_destino: N(2)
		m_credito: N(10)
	FinRegistro
	viaje: archivo de VIAJES 
	reg_v: VIAJES 
	B: arreglo[1...13,1...24,1...2] de enteros 
	i, j, k, menor, prom, d_mas, d_menor, m_menor, mas: entero 
	des: AN
	Funcion obtener_destino(j:entero): AN 
		des:= j 
		Para i:= 1 a 23 hacer 
	 	 	Si A[i] = des entonces
				obtener_destino:= A[i] 
			FinSi
		FinPara
	FinFuncion
 	Proceso 
		Abrir E/(viaje)
		Leer(viaje,reg_v)
		menor:=HV; prom:=0; d_mas:=0; d_menor:=0; m_menor:=0; mas:=0
		Para k:= 1 a 2 hacer 
			Para i:= 1 a 13 hacer 
				Para i 1 a 24 hacer 
					B[i,j,k]:= 0 
				FinPara
			FinPara 
		FinPara
		i:=0; j:=0; k:=0
		Mientras NFDA(viaje) hacer 
			k:=1
			i:= reg_v.fecha.mes 
			j:= reg_v.p_destino 
			B[i,j,k]:= B[i,j,k] + reg_v.m_credito
			B[13,j,k]:= B[13,j,k] + reg_v.m_credito
			B[i,24,k]:=B[i,24,k] + reg_v.m_credito
			B[13,24,1]:= B[13,24,1] + reg_v.m_credito
			k:= 2 
			B[i,j,k]:= B[i,j,k] + 1 
			B[13,j,k]:= B[13,j,k] + 1
			B[i,24,k]:= B[i,24,k] + 1
			Leer(viaje,reg_v)
		FinMientras
		Para k:= 1 a 2 hacer 
			Para i:= 1 a 13 hacer 
				Para j:= 1 a 24 hacer 
			 	 	Si B[i,j,1] < menor entonces 
						menor:= B[i,j,1]
						d_menor:= j
						m_menor:= i 
					FinSi

					Esc("Montos por destino",obtener_destino(j),"en el mes",i, B[i,j,1],"$") // consigna 1
					Esc("Cantidad de viajes por destino",obtener_destino(j),"en el mes",i, B[i,j,2]) // consigna 1

				 	Si B[11,j,2] > d_mas entonces 
						d_mas:= B[11,j,2]
						mas:= j
					FinSi

					Si i <> 13 entonces 
						prom:= (B[i,24,1]/B[13,24,1])
						Esc("Promedio en el mes",i,"es de",prom) // consigna 3
					Fsi 

				FinPara
			FinPara 
		FinPara 
		j:= mas 
		Esc("Menor monto de: ",menor,"$"," en el mes ",m_menor," del destino ",obtener_destino(d_menor)) // consigna 2
		Esc("El destino mas elegido en noviembre es",obtener_destino(j)) // consigna 4
		Cerrar(viaje)
	FinProceso
FinAccion


 



	 	 
  
  

			
						 
				 
		 	 
	 
 