Accion Ejercicio1 es 
 Ambiente 
 Fecha = registro 
	 dia: N(2)
	 mes: N(2)
	 anio: N(4)
	FinRegistro
 CLIENTES = registro 
	 clave = registro 
	  	 nro_cliente: N(5)
		FinRegistro
	 apeynom: AN(30)
	 dni: N(8)
	 id_paquete: N(1)
	 saldo: N(10)
	 estado: AN(12)
	 cat: AN(10)
	 puntos: N(4)
	 f_baja: Fecha
	FinRegistro
 arch, arch_act: archivo de CLIENTES ordenado por clave 
 reg, reg_act, aux: CLIENTES
 NOVEDADES = registro 
	 clave = registro 
	 	 nro_cliente: N(5)
		 nro_novedad: (0...999)
		FinRegistro
	 apeynom: AN(30)
	 dni: N(8)
	 id_paquete: N(1)
	 f_novedad: Fecha
	 monto: N(10)
	FinRegistro
 arch_nov: archivo de NOVEDADES ordenado por clave 
 reg_nov: NOVEDADES
 PAQUETES_TURISTICOS = registro 
	 id_paquete: N(1)
	 nombre: AN(10)
	 costo: N(10)
	 destino: AN(15)
	FinRegistro
 arch_turistico: archivo de PAQUETES_TURISTICOS indexado por id_paquete
 reg_turistico: PAQUETES_TURISTICOS
 m_total, c_baja, c_simple, c_plata, c_oro, c_diamante, total: entero
 Procedimiento leer_nov() es 
	 Leer(arch_nov,reg_nov)
	 	Si FDA(arch_nov) entonces 
		 reg_nov.clave:= HV
		FinSi 
	FinProcedimiento 
 Procedimiento leer_cli() es 
	 Leer(arch,reg)
	 	Si FDA(arch) entonces 
		 reg.clave:= HV
		FinSi
	FinProcedimiento
 Procedimiento contar_cat() es 
	 Segun reg.cat hacer 
	 	 "simple": c_simple:= c_simple + 1 
		 "plata": c_plata:= c_plata + 1 
		 "oro": c_oro:= c_oro + 1 
		 "diamante": c_diamante:= c_diamante + 1
		FinSegun 
	FinProcedimiento
 Procedimiento procesos_iguales() es 
	 	Si reg_nov.nro_novedad = 0 entonces 
		 Esc("Alta no disponible")
		Sino 
			Si reg_nov.nro_novedad = 999 entonces 
			 c_baja:= c_baja + 1
			 m_total:= m_total + reg_nov.monto 
			 aux.f_baja:= reg_nov.f_novedad
			Sino 
			 Esc("Pago del cliente",reg_nov.apeynom)
			FinSi
		FinSi
	FinProcedimiento
 Procedimiento porcentaje() es 
	 c_diamante:= (total/c_diamante)*100; Esc("Porcentaje diamante",c_diamante,"%")
	 c_oro:= (total/c_oro)*100; Esc("Porcentaje oro",c_oro,"%")
	 c_plata:= (total/c_plata)*100; Esc("Porcentaje plata",c_plata,"%")
	 c_simple:= (total/c_simple)*100; Esc("Porcentaje simple",c_simple,"%")
	FinProcedimiento
 Proceso 
	 Abrir E/(arch)
	 Abrir E/(arch_nov)
	 Abrir E/(arch_turistico)
	 leer_cli()
	 leer_nov()
	 c_baja:= 0; c_diamante:=0; c_oro:=0; c_plata:=0; c_simple:=0, m_total:=0; total:=0
	 Mientras (reg.clave <> HV) o (reg_nov.clave <> HV) hacer 
	 	 	Si reg.clave < reg_nov.clave entonces 
			 reg_act:= reg
			 contar_cat()
			 Grabar(arch_act,reg_act)
			Sino 
				Si reg.clave = reg_nov.clave entonces 
				 aux:= reg 
				 Mientras aux.clave = reg_nov.clave hacer 
				 	 procesos_iguales()
					 leer_nov()
					FinMientras
				 contar_cat() 
				 reg_act:= aux 
				 Grabar(arch_act,reg_act)
				 leer_cli()
				Sino 
					Si reg_nov.nro_novedad = 0 entonces 
					 aux.nro_cliente:= reg_nov.nro_cliente
					 aux.apeynom:= reg_nov.apeynom
					 aux.dni:= reg_nov.dni 
					 aux.id_paquete:= reg_nov.id_paquete
					 aux.estado:= "deudor"
					 aux.cat:= "simple"
					 aux.puntos:= 0 
					 aux.f_baja:= " "
					FinSi
				 leer_nov()
				 contar_cat()
				 Mientras aux.clave = reg_nov.clave hacer 
				 	 procesos_iguales()
					 leer_nov()
					FinMientras
				 reg_act:= aux 
				 Grabar(arch_act,reg_act)
				FinSi
			FinSi 
		FinMientras
	 total:= c_diamante + c_oro + c_plata + c_simple 
	 porcentaje()
	 Esc("La cantidad de usuarios dados de baja es",c_baja,"monto total a reintegrar",m_total)
	 Cerrar(arch)
	 Cerrar(arch_nov)
	 Cerrar(arch_act)
	 Cerrar(arch_turistico)
	FinProceso 
FinAccion

Accion Ejercicio2 es 
 Ambiente
 Fecha = registro 
	 dia: N(2)
	 mes: N(2)
	 anio: N(4)
	FinRegistro
 CLIENTES = registro 
  	 clave = registro 
		  nro_cliente: N(5)
	   FinRegistro
	 apeynom: AN(30)
	 dni: N(8)
	 id_paquete: N(1)
	 saldo: N(10)
	 estado: AN(12)
	 cat: AN(10)
	 puntos: N(4)
	 f_baja: Fecha
    FinRegistro
 arch: archivo de CLIENTES ordenado por clave 
 reg: CLIENTES
 PAQUETES_TURISTICOS = registro 	
  	 id_paquete: N(1)
	 nombre: AN(10)
	 costo: N(10)
	 destino: AN(15)
	FinRegistro
 arch_turistico: archivo de PAQUETES_TURISTICOS indexado por id_paquete
 reg_turistico: PAQUETES_TURISTICOS
 clientes: arreglo[1...5,1...13] de enteros 
 mayor, paquete_may: entero 
 Funcion obtener_paquete(x: entero): AN 
	 reg_turistico.id_paquete:= x 
	 Leer(arch_turistico,reg_turistico)
	 	Si EXISTE entonces 
		 Esc(reg_turistico.nombre)
		FinSi
	FinFuncion
 Proceso
	 Abrir E/(arch)
	 Abrir E/(arch_turistico)
	 Leer(arch,reg)
	 mayor:=0; paquete_may:=0
	 Para i:= 1 a 5 hacer 
	 	 Para j:= 1 a 13 hacer 
		 	 clientes[i,j]:=0
			FinPara
		FinPara
	 Mientras NDFA(arch) hacer 
	 		Si (reg.estado = "saldado") o (reg.estado = "saldo a favor") entonces 
	 		 segun reg.cat hacer 
			 	 "simple": i:= 1 
				 "plata": i:= 2 
				 "oro": i:= 3
				 "diamante": i:= 4 
				FinSegun 
			 j:= reg.id_paquete
			 clientes[i,j]:= clientes[i,j] + 1 
			 clientes[5,j]:= clientes[5,j] + 1 
			 clientes[i,13]:= clientes[i,13] + 1 
			 clientes[5,13]:= clientes[5,13] + 1
			FinSi
		 Leer(arch,reg)
		FinMientras
	 Para i:= 1 a 5 hacer 
	 	 Para j:= 1 a 12 hacer 
		 	 	Si clientes[5,j] > mayor entonces 
				 mayor:= clientes[5,j]
				 paquete_may:= j 
				FinSi
			FinPara
		FinPara
	 Esc("El total de paquetes vendidos es de",clientes[5,13])
	 Esc("El paquete mas vendido es",obtener_paquete(paquete_may))
	 Cerrar(arch)
	 Cerrar(arch_turistico)
	FinProceso
FinAccion

		 
		

					 

 
