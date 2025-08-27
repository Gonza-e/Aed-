Accion Ejercicio1 es 
 Ambiente 
 Fecha = registro 
	 dia: N(2)
	 mes: N(2)
	 anio: N(4)
	FinRegistro 
 INTERCAMBIOS = registro 
	 cod_figurita: N(5)
	 cod_amigo: N(5)
	 f_solicitud: Fecha
	FinRegistro
 ALBUM = registro 
	 cod_figurita: N(5)
	 cantidad: N(2)
	 p_duplicados: ("Si","No")
	FinRegistro 
 alb, salida: archivos de ALBUM ordenado por cod_figurita 
 reg_a, aux: ALBUM
 inter: archivo de INTERCAMBIOS ordenado por cod_figurita 
 reg_a: INTERCAMBIOS
 c_duplicadas: entero 
 Procedimiento leer_alb() es 
	 Leer(alb, reg_a)
	 	Si FDA(alb) entonces
		 reg_a.cod_figurita:= HV
		FinSi
	FinProcedimiento
 Procedimiento leer_inter() es 
	 Leer(inter_ reg_i) 
	 	Si FDA(inter) entonces 
		 reg_a.cod_figurita:= HV
		FinSi 
	FinProcedimiento 
 Procedimiento procesos_iguales() es 
		Si reg_a.p_duplicados = "Si" entonces 
		 aux.cantidad:= aux.cantidad + 1 
		Sino 
		 Esc("La figurita no acepta duplicado")
		FinSi
	FinProcedimiento
 Proceso 
	 Abrir E/(alb)
	 Abrir E/(inter)
	 Abrir S/(salida)
	 leer_alb()
	 leer_inter()
	 c_duplicadas:= 0 
	 Mientras (reg_a.cod_figurita <> HV) o (reg_i.cod_figurita <> HV) hacer 
	 	 	Si reg_a.cod_figurita < reg_i.cod_figurita entonces
			 sal:= reg_a
			Sino 
				Si reg_a.cod_figurita = reg_i.cod_figurita entonces
				 aux:= reg_a 
				 Mientras aux.cod_figurita = reg_i.cod_figurita hacer
				 	 procesos_iguales()
					 c_duplicadas:= c_duplicadas + 1 
					FinMientras
				 sal:= aux 
				 Grabar(salida, sal)
				 leer_inter()
				Sino 
				 	Si reg_a.cod_figurita > reg_i.cod_figurita entonces 
					 	Si diff_dias(7,reg_i.f_solicitud) entonces
						 sal.cod_figurita:= reg_i.cod_figurita
						 sal.cantidad:= 1 
						 sal.p_duplicados:= "No"
						 Grabar(salida, sal)
						Sino 
						 aux:= reg_a 
						 Mientras aux.cod_figurita = reg_i.cod_figurita hacer 
						 	 leer_inter()
							FinMientras
						FinSi
					 leer_alb()
					FinSi
				FinSi
			FinSi
		FinMientras
	 Cerrar(alb)
	 Cerrar(inter)
	 Cerrar(salida)
	 Esc("La cantidad de figuritas duplicadas es de", c_duplicadas)
	FinProceso
FinAccion

Accion Ejercicio2 es 
 Ambiente 
	ALBUM = registro 
		cod_usuario: N(5)
		cod_figurita: N(5)
		cantidad: N(2)
		tipo: AN(1)
	FinRegistro
	alb: archivo de ALBUM ordenado por cod_usuario
	reg_alb: ALBUM

	AMIGOS = registro 
		cod_usuario: N(5)
		apellido: AN(15)
		nombre: AN(15)
		celular: N(10)
	FinRegistro
	ami: archivo de AMIGOS indexado por cod_usuario
	reg_ami: AMIGOS

	A: arreglo[1...4,1...11] de ALBUM
	porc, t_usuario, i, j, mayor, res_cod: entero   
	ape, nom: AN

	Procedimiento obtener_album() es 
		Segun i hacer 
			1: "Dorado"
			2: "Comun"
			3: "Virtual"
		FinSegun
	FinProcedimiento
	
	Proceso 
		Abrir E/(alb)
		Abrir E/(ami)
		Leer(alb,reg_alb)
		porc:=0, t_usuario:=0; mayor:=0
		i:=0; j:=1
		res_cod:= reg_alb.cod_usuario
		Para i:= 1 a 4 hacer 
			Para j:= 1 a 10 hacer 
				A[i]:=0
			FinPara 
		FinPara  
	 Mientras NFDA(alb) hacer 
	 	 	Si res_cod <> reg_alb.cod_usuario entonces
			 	Si t_usuario > mayor entonces 
					mayor:= t_usuario 
					reg_ami.cod_usuario:= res_cod
					Leer(ami,reg_ami)
				 	Si EXISTE entonces 
						ape:= reg_ami.apellido
						nom:= reg_ami.nombre 
					FinSi
				FinSi
				j:= j+1
				t_usuario:=0 
				res_cod:= reg_alb.cod_usuario  
			FinSi

			Segun reg_alb.tipo hacer 
				"D": i:=1
				"C": i:=2
				"V": i:=3
			FinSegun 

			A[i,j]:= A[i,j] + reg_alb.cantidad
			A[4,j]:= A[4,j] + reg_alb.cantidad
			A[i,11]:= A[i,11] + reg_alb.cantidad
			A[4,11]:= A[4,11] + reg_alb.cantidad
			t_usuario:= t_usuario + reg_alb.cantidad
			Leer(alb,reg_alb)
		FinMientras

		Para i:= 1 a 4 hacer 
			Para j:= 1 a 11 hacer 
				Esc("El usuario",j,"cuenta con",A[4,j],"figuritas")
				Esc("El album",obtener_album(),"cuenta con",A[i,11])
				porc:= (A[4,j]/A[4,11])*100
				Esc("El porcentaje de figuritas del usuario",j,"es de",porc,"%")
			FinPara
  	    FinPara
		Esc("El usuario con mayor cantidad de figuritas es", ape, nom)
		Cerrar(ami)
		Cerrar(alb)	 
	FinProceso 
FinAccion
		

				 
			 

			