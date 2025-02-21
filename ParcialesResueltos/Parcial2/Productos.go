Accion Ejercicio1 (jugador: arreglo[1...26] de AN) es 
 Ambiente 
 STOCK = registro 
	 clave = registro 
	 	 id_producto: N(5)
		FinRegistro 
	 nombre: AN(30)
	 descripcion: AN(100)
	 cat: AN(1)
	 new_lanzamiento: ("Si","No")
	 porc_desc: N(2,2)
	 stock: N(3)
	FinRegistro
 arch_stock, stock_act: archivo de STOCK ordenado por clave 
 reg_stock, reg_act, aux: STOCK
 PREVENTA = registro 
	 clave = registro 
	 	 id_producto: N(5)
		 id_cliente: N(5)
		FinRegistro
	 cantidad: N(3)
	 es_personzalidado: AN(2)
	 nro_jugador: (1...26)
	 nom_jugador: AN(30)
	 talle: AN(1)
	FinRegistro
 arch_pre: archivo de PREVENTA ordenado por clave 
 reg_pre: PREVENTA
 A: arreglo[1...26] de enteros 
 cat: arreglo[1...4] de enteros
 i, c_sinstock, calc_stock, menor, cat_menor: entero
 Procedimiento leer_stock() es 
	 Leer(arch_stock,reg_stock)
	 	Si FDA entonces 
		 reg_stock.clave:= HV
		FinSi 
	FinProcedimiento 
 Procedimiento leer_preventa() es 
	 Leer(arch_pre,reg_pre)
	 	Si FDA entonces 
		 reg_pre.clave:= HV
		FinSi
	FinProcedimiento
 Procedimiento contar_cat() es 
	 segun reg_stock.cat hacer 
	 	 "C": i:= 1 
		 "R": i:= 2 
		 "G": i:= 3 
		 "S": i:= 4
		FinSegun 
	 A[i]:= A[i] + 1
	FinProcedimiento
 Procedimiento procesos_iguales() es 
	 Mientras aux.clave = reg_pre.clave hacer 
		 	Si reg_stock.stock > reg_pre.cantidad entonces
			 aux.stock:= aux.stock - reg_pre.cantidad
			FinSi
		 calc_stock:= (reg_stock.stock - reg_pre.cantidad)
			Si calc_stock <= 0 entonces 
			 calc_stock:= calc_stock * (-1)
			FinSi
		 c_sinstock:= c_sinstock + calc_stock
		 i:= reg_pre.nro_jugador
		 A[i]:= A[i] + 1 
		 contar_cat()
		 leer_preventa()
		FinMientras
	FinProcedimiento 
 Proceso 
	 Abrir E/(arch_pre)
	 Abrir E/(arch_stock)
	 Abrir /S(stock_act)
	 leer_preventa()
	 leer_stock()
	 i:=0; c_sinstock:=0; calc_stock:=0; menor:=HV; cat_menor:=0; jug_mayor:=0; mayor:=0
	 Para i:= 1 a 26 hacer 
	 	 A[i]:= 0 
		FinPara 
	 Para i:= 1 a 4 hacer 
	 	 cat[i]:= 0 
		FinPara
	 Mientras (reg_stock.clave <> HV) o (reg_pre.clave <> HV) hacer 
	 	 	Si reg_stock.clave < reg_pre.clave entonces 
			 reg_act:= reg_stock
			 Grabar(stock_act,reg_act)
			 contar_cat()
			 leer_stock()
			Sino 
				Si reg_stock.clave = reg_pre.clave entonces 
				 aux:= reg_stock 
				 procesos_iguales()
				 reg_act:= aux 
				 Grabar(stock_act,reg_act)
				 leer_stock()
				Sino 
				 c_sinstock:= c_sinstock + reg_pre.cantidad
				 i:= reg_pre.nro_jugador
				 A[i]:= A[i] + 1 
				 leer_preventa()
				 procesos_iguales()
				FinSi
			FinSi
		FinMientras
	 Para i:= 1 a 26 hacer 
	 	 	Si A[i] > mayor entonces 
			 mayor:= A[i]
			 jug_mayor:= i 
			FinSi 
		FinPara 
	 Para i:= 1 a 4 hacer 
	 	 	Si cat[i] < menor entonces 
			 menor:= cat[i]
			 cat_menor:= i 
			FinSi
		FinPara 
	 Esc("La cantidad de productos sin stock es",c_sinstock)
	 Esc("El jugador con mas camisetas vendidas es",jugador[jug_mayor])
	 Esc("La categoria menos comprada es",obtener_cat(cat_menor))
	 Cerrar(arch_pre)
	 Cerrar(arch_stock)
	 Cerrar(stock_act)
	FinProceso 
FinAccion

Accion Ejercicio2 es 
 Ambiente 
 PROMOCIONES = registro 
	 cod_suc: (1...10)
	 cod_promocion: (0...5)
	 cod_prod: AN(7)
	 cant: N(4)
	FinRegistro
 arch: archivo de PROMOCIONES
 reg: PROMOCIONES
 promos: arreglo[1...7,1...11] de Datos 
 Datos = registro 
	 importe: N(6)
	 cant: N(4)
	FinRegistro 
 promedio, mayor, prom_may, mayor2, monto: entero 
 Proceso 
	 Abrir E/(arch)
	 Leer(arch,reg)
	 promedio:=0; mayor:=0; prom_may:=0; mayor2:=0; monto:=0; suc_may:=0
	 Para i:= 1 a 7 hacer 
	 	 Para j:= 1 a 11 hacer 
		 	 promos[i,j]:=0 
			FinPara
		FinPara
	 Mientras NFDA(arch) hacer 
	 	 i:= reg.cod_promocion
		 j:= reg.cod_suc
		 	Si i <> 0 entonces 
			 promos[i,j].importe:= promos[i,j].importe + getImporte(reg.cod_prod)
			 promos[i,j].cant:= promos[i,j].cant + reg.cant
			 promos[7,j].importe:= promos[7,j].importe + getImporte(reg.cod_prod)
			 promos[7,j].cant:= promos[7,j].cant + reg.cant
			 promos[i,11].importe:= promos[i,11].importe + getImporte(reg.cod_prod)
			 promos[i,11].cant:= promos[i,11].cant + reg.cant
			 promos[7,11].importe:= promos[7,11].importe + getImporte(reg.cod_prod)
			 promos[7,11].cant:= promos[7,11].cant + reg.cant
			FinSi
		 Leer(arch,reg)
		FinMientras
	 Para i:= 2 a 6 hacer
	 	 Para j:= 1 a 11 hacer 
		 	 	Si promos[i,11].importe > mayor entonces 
				 mayor:= promos[i,11].importe
				 prom_may:= i 
				FinSi
				Si j <> 11 entonces
					Si promos[i,j]importe > mayor2 entonces
					 mayor2:= promos[i,j].importe
					 suc_may:= j 
					FinSi
				FinSi
			FinPara
		 Esc("Para la promocion",i,"la sucursal con mayor importe fue",suc_may)
		 suc_may:= 0 
		 mayor2:= 0
		FinPara
	 Esc("La promocion que mayor importe recaudo es:",prom_may)
	 Cerrar(arch)
	FinProceso
FinAccion
		 
					    
			  
			  	 
				 


				 

