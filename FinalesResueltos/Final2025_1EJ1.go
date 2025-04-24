Accion Ejercicio1 es 
 Ambiente 
	VIAJES = registro 
		nro_viaje: N(3)
		ciu_origen: (1...5)
		ciu_destino: (1...5)
		p_disponibles: N(2)
	FReg 

	arch: archivo de VIAJES 
	reg: VIAJES 

	nodo = registro 
		datos: VIAJES 	
		vendidos: N(3)
		tipo: Arreglo[1...2] de enteros 
		prox: puntero a nodo 
	FReg 

	p,prim,q,s: puntero a nodo 

	cant_tipo: Arreglo[1...2] de enteros 
	porc: real 
	no_50, total_vend, total, mayor, num_viaje, USU_origen, USU_destino: entero
	tipo, opcion: AN(2)

	procedimiento sumar_descontarViaje() es 
		si (tipo = "E") y (*p.tipo <> 0) entonces 
			*p.tipo[1]:= *p.tipo[1] - 1 
			cant_tipo[1]:= cant_tipo[1] + 1  
			*p.vendidos:= *p.vendidos + 1 
		sino 
			si (tipo = "V") y (*p.tipo <> 0) entonces 
				*p.tipo[2]:= *p.tipo[2] - 1 
				cant_tipo[2]:= cant_tipo[2] + 1
				*p.vendidos:= *p.vendidos + 1 
			sino 
				Esc("La reserva no se puede hacer")
			fsi  
		fsi 
	FP 

	procedimiento cargar_tipos() es 
		*q.tipo[1]:= *q.datos.p_disponibles * 0.3 
		*q.tipo[2]:= *q.datos.p_disponibles * 0.7 
	FP 

 Proceso 
	 	Abrir E/(arch) 
	 	Leer(arch,reg) 
	 	no_50:= 0; total_vend:= 0; total:= 0; mayor:= 0; num_viaje:= 0; tipo:= " "
		prim:= nil

	 	Mientras NFDA(arch) hacer 
	 	  	Nuevo(q); *q.datos:= reg; *q.vendidos:= 0; cargar_tipos()
			*q.prox:= nil; prim:= q
			total:= total + 1  
			Leer(arch,reg)
		FM 

		Repetir
			Esc("Ingrese origen, destino y tipo de asiento")
			Esc("E: economico, V: vip")
			Leer(USU_ciudad,USU_destino,tipo)

			p:= prim 
			Mientras (p<>nil) y ((*p.datos.ciu_origen <> USU_origen) y (*p.datos.ciu_destino <> USU_destino)) hacer
				p:= *p.prox 
			FM 
			
			Si (*p.datos.ciu_origen = USU_origen) y (*p.datos.ciu_destino = USU_destino) entonces 
				sumar_descontarViaje()
			sino 
				Esc("El viaje no se encontro")
			fsi 
			
			Esc("Desea continuar? si o no")
			Leer(opcion)		
		Hasta que (opcion = no)

		p:= prim 
		Mientras (p<>nil) hacer 
			total_vend:= total_vend + *p.vendidos 

			si *p.datos.p_disponibles < *p.vendidos * 0.5 entonces 
				no_50:= no_50 + 1 
			fsi 

			si mayor < *p.vendidos entonces 
				mayor:= *p.vendidos 
				num_viaje:= *p.datos.nro_viaje
			fsi 

			p:= *p.prox 
		FM 

		porc:= (no_50/total)*100
		Esc("Cantidad de asientos vip vendidos: ",cant_tipo[2]," Cantidad de asientos economicos vendidos: ",cant_tipo[1]) //Consigna 1
		Esc("El numero de viaje con mas boletos vendidos es de: ",num_viaje) //Consigna 2 
		Esc("El total de boletos vendidos es de: ",total_vend) //Consigna 4 
		Esc("Porcentaje de viajes que no cubren el 50%: ",porc:0:2,"%") // Consigna 3 
		Cerrar(arch)
	FProceso 
FAccion

		

