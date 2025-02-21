Accion Ejercicio1 (tipoventa: arreglo[1...10] de AN) es 
 Ambiente 
 VENTAS = registro 
	 nro_venta: N(6)
	 nro_pro: N(6)
	 barco: (1,2,3)
	 t_venta: N(2)
	 estado: ("D","FT")
	FinRegistro 
 vent: archivo de VENTAS ordenado por nro_venta
 reg_v: VENTAS 
 PRODUCTOS = registro
	 barco: (1,2,3)
	 nro_pro: N(6)
	 tipo: ("I","P")
	 t_venta: N(2)
	 fumador: ("SI","NO")
	 cat: (1,2,3)
	FinRegistro
 prod: archivo de PRODUCTOS indexado por nro_pro y barco 
 reg_p: PRODUCTOS
 venta2y5, tipoFT, porcentaje: entero 
 Procedimiento mod_categoria() es 
	 reg_p.cat:= 3 
	 Regrabar(prod,reg_p)
	FinProcedimiento
 Proceso 
	 Abrir E/S (prod)
	 Abrir E/(vent)
	 Leer(vent_reg_v)
	 venta2y5:=0; tipoFT:=0, porcentaje:=0
	 Mientras NFDA(vent) hacer 
	 	 reg_p.nro_pro:= reg_v.nro_pro 
		 Leer(prod,reg_p)
		 	Si EXISTE entonces 
			 	Si (reg_p.t_venta = 2) O (reg_p.t_venta = 5)
				 venta2y5:= venta2y5 + 1 
				 	Si reg_v.estado = "FT" entonces
					 tipoFT:= tipoFT + 1 
					 Borrar(prod,reg_p)
					FinSi 
				FinSi
				Si reg_v.t_venta = 10 entonces
				 reg_p.cat:= 3 
				 Regrabar(prod,reg_p)
				FinSi
			FinSi
		 Leer(vent,reg_v)
		FinMientras 
	 porcentaje:= (tipoFT/venta2y5)*100
	 Esc("La cantidad de ventas tipo 2 y 5 son de:",venta2y5)
	 Esc("El porcentaje de ventas FT es de:",porcentaje,"%")
	 Cerrar(vent) 
	 Cerrar(prod)
	FinProceso 
FinAccion

Accion Ejercicio2 es 
 Ambiente
 Fecha = registro 
	 dia: N(2)
	 mes: N(2)
	 anio: N(4)
	FinRegistro
 TOTALES = registro 
	 barco: (1,2,3)
	 fecha: Fecha
	 tipo: ("I","P")
	 importe: N(7,2)
	FinRegistro
 tot: archivo de TOTALES 
 reg_t: TOTALES
 A: arreglo[1...3,1...13,1...3] de reales
 mayor, menos, mes_may, mes_men, bar, tip, k, i, j: entero
 Funcion obtener_tipo(i:entero): AN 
	 Segun i hacer 
	 	 1: obtener_tipo:= "Indumentaria"
		 2: obtener_tipo:= "Perfumeria"
		FinSegun
	FinFuncion
 Proceso 
	 Abrir E/(tot)
	 Leer(tot,reg_t)
	 mayor:=0; menos:=HV; mes_may:=0; mes_men:=0; bar:=0, tip:=0
	 Para k:= 1 a 3 hacer 
	 	 Para i:= 1 a 3 hacer 
		 	 Para j:= 1 a 13 hacer 
			 	 A[i,j,k]:=0
				FinPara
			FinPara 
		FinPara 
	 Mientras NFDA(tot) hacer 
	 	 k:= reg_t.barco
		 j:= reg_t.fecha.mes 
		 i:= 1 
		 	Si reg_t.tipo = "P" entonces 
			 i:= 2 
			FinSi
		 A[i,j,k]:= A[i,j,k] + reg_t.importe
		 A[3,j,k]:= A[3,j,k] + reg_t.importe
		 A[i,13,k]:= A[i,13,k] + reg_t.importe
		 A[3,13,k]:= A[3,13,k] + reg_t.importe
		 Leer(tot,reg_t)
		FinMientras
	 Para k:= 1 a 3 hacer 
	 	 Para i:= 1 a 3 hacer 
			 Para j:= 1 a 13 hacer 
			 	  	Si A[i,13,k] > mayor entonces 
					 mayor:= A[i,13,k]
					 mes_may:= j 
					 bar:= k 
					 tip:= i
					FinSi
				 	Si A[3,j,k] < menos entonces 
					 menos:= A[3,j,k] 
					 mes_men:= j 
					FinSi
				FinSi
			FinSi
		 Esc("En el barco",k,"en el mes",mes_men,"se recaudo menos")
		 menos:= HV
		FinPara 
	 i:= tip
	 Esc("El tipo con mayor importe es",obtener_tipo(i),"en el mes",mes_may,"en el barco",bar)
	 Cerrar(tot)
	FinProceso
FinAccion


	 


				 