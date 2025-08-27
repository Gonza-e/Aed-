Accion Ejercicio1 es 
 Ambiente 
	Fecha = registro 
		dia: N(2)
		mes: N(2)
		anio: N(4)
	FinRegistro
	BICICLETAS = registro 
		clave = registro 
			nro_serie: N(4)
			modelo: AN(10)
		FinRegistro 
		f_adq: Fecha
		f_ult_mant: Fecha
		disp: (verdadero, falso)
	FinRegistro
	bici, salida: archivos de BICICLETAS ordenado por clave 
	reg_b, sal, aux: BICICLETAS
	NOVEDADES = registro 
		clave = registro	
			nro_serie: N(4)
			modelo: AN(10)
			tipo_novedad: N(1)
			f_novedad: Fecha
		FinRegistro
		hora_i: N(4)
		hora_f: N(4)
		nro_circuito: N(5)
		id_usu: N(5)
	FinRegistro
	nov: archivo de NOVEDADES ordenado por clave 
	reg_n: NOVEDADES
	USUARIOS = registro 
		id_usu: N(5)
		DNI: N(8)
		sex: (m o f)
		apeynom: AN(30)
		domicilio: AN(20)
		localidad: AN(20)
		provincia: AN(20)
		edad: N(2)
	FinRegistro
	usu: archivo de USUARIOS indexado por id_usu 
	reg_u: USUARIOS
	p_hombre, p_mujer, total: entero

	Procedimiento procesos_iguales() es 
	 	Si reg_n.clave.tipo_novedad = 1 entonces 
		 	Esc("Error")
		Sino 
			Si reg_n.clave.tipo_novedad = 2 entonces 
				reg_u.id_usu:= reg_n.id_usu
				Leer(usu, reg_u) 
			 	Si EXISTE entonces
				 	Segun reg_u.sexo hacer 
						M: p_hombre:= p_hombre + 1 
						F: p_mujer:= p_mujer + 1
					FinSegun 
				Sino 
				 	Esc("El usuario no existe")
				FinSi 
			Sino 
				aux.nro_serie:= reg_n.nro_serie
				aux.modelo:= reg_n.modelo
				aux.f_ult_mant:= reg_n.f_novedad
				aux.disp:= falso 
			FinSi
		FinSi
	FinProcedimiento 

	Procedimiento leer_bici() es 
	 	Leer(bici, reg_b)
	 	Si FDA(bici) entonces
		 	reg_b.clave:= HV
		FinSi
	FinProcedimiento

	Procedimiento leer_nov() es 
	 	Leer(nov, reg_n)
	 	Si FDA(nov) entonces
		 	reg_n.clave:= HV 
		FinSi
	FinProcedimiento

	Proceso 
		Abrir E/(bici)
		Abrir E/(nov)
		Abrir E/(usu)
		Abrir S/(salida)
		leer_bici()
		leer_nov()
		p_hombre:= 0; p_mujer:= 0; total:= 0
		Mientras (reg_b.clave <> HV) o (reg_n.clave <> HV) hacer 
	 		Si reg_b.clave < reg_n.clave entonces 
				sal:= reg_b 
				Grabar(salida, sal)
				leer_bici()
			Sino 
				Si reg_b.clave = reg_n.clave entonces
				 	aux:= reg_b
					Mientras aux.clave = reg_n.clave hacer 
						procesos_iguales()
						leer_nov()
					FinMientras 
					sal:= aux 
					Grabar(salida, sal)
					leer_bici()
				Sino 
					Si reg_b.clave > reg_n.clave entonces 
						Si reg_n.clave.tipo_novedad = 1 entonces
							aux.clave.nro_serie:= reg_n.nro_serie
							aux.clave.modelo:= reg_n.modelo
							aux.f_ult_mant:= " "
							aux.disp:= verdadero
						FinSi
						leer_nov() 
						Mientras aux.clave = reg_b.clave entonces 
							procesos_iguales()
							leer_nov()
						FinMientras
						sal:= aux 
						Grabar(salida, sal)
					FinSi
				FinSi
			FinSi
		FinMientras
		Cerrar(nov)
		Cerrar(bici)
		Cerrar(salida)
		Cerrar(usu)
		Esc("La cantidad de prestamos por hombre fue",p_hombre,"y por mujer",p_mujer)
	FinProceso 
FinAccion

		      
Accion Ejercicio2 (A: arreglo[1...2,1...6] de enteros) es 
 Ambiente 
	 NOVEDADES = registro 
		clave = registro	
			nro_serie: N(4)
			modelo: AN(10)
			tipo_novedad: N(1)
			f_novedad: Fecha
		FinRegistro
		hora_i: N(4)
		hora_f: N(4)
		nro_circuito: N(5)
		id_usu: N(5)
	FinRegistro
	B: arreglo[1...3,1...7,1...2] de enteros 
	nov: archivo de NOVEDADES ordenado por clave 
	reg_n: NOVEDADES
	paseo, circuito, mayor, i, j, k, valor, c_may, c_men: entero
	Proceso 
		Abrir E/(nov)
		Leer(nov,reg_n)
		paseo:=0; circuito:=0; mayor:=LV; valor:=0; c_may:=0; c_men:=0; prestamos:= 0 
		Para k:= 1 a 2 hacer 
			Para i:= 1 a 2 hacer 
				Para j:= 1 a 3 hacer 
					B[k,i,j]:=0 
				FinPara
			FinPara
		FinPara 
		k:=0; i:=0; j:=0
		Esc("Ingrese el circuito y tipo de paseo")
		Leer(circuito); Leer(paseo)
		Mientras NFDA(nov) hacer 
	 		Si reg_n.tipo_novedad = 2 entonces
				prestamos:= prestamos + 1 
				i:= 1 
				k:= 1 
				j:= reg_n.nro_circuito = //5 
				valor:= 1500
			 	Si diff_horas(reg_n.hora_i,reg_n.hora_f) < 6 entonces
					i:= 2
					valor:= 1000 //NO ENTRO
				FinSi
				B[k,i,j]:= B[k,i,j] + (valor + A[i,j])
				B[k,3,j]:= B[k,3,j] + (valor + A[i,j])
				B[k,i,7]:= B[k,i,7] + (valor + A[i,j])
				B[k,3,7]:= B[k,3,7] + (valor + A[i,j])
				k:= 2 
				B[k,i,j]:= B[k,i,j] + 1
				B[k,3,j]:= B[k,3,j] + 1 
				B[k,i,7]:= B[k,i,7] + 1 
				B[k,3,7]:= B[k,3,7] + 1 
			FinSi
		 	Leer(nov,reg_n)
		FinMientras 
		Para k:= 1 a 2 hacer 
			Para i:= 1 a 3 hacer 
				Para j:= 1 a 7 hacer
			 	  	Si B[2,i,j] > mayor entonces
						mayor:= B[2,i,j] 
						c_may:= j
					FinSi
				FinPara
			FinPara
		FinPara
		Esc("El circuito con mayor cantidad de paseos es",c_may)
		i:= paseo 
		j:= circuito 
		Esc("Total por circuito y paseo ingresado por el usuario"); Esc("Paseo y circuito: $", B[1,i,j])
		Esc("Importe total: $", B[1,3,6], "Total de paseos:", B[2,3,6])
		Esc("La cantidad de prestamos es de:",prestamos)
		Cerrar(nov)
	FinProceso 
FinAccion
