accion 2025 es 
 Ambiente 
	Fecha = registro 
		dia: N(2)
		mes: N(2)
		año: N(4)
	FReg

	EJEMPLARES = registro 
		clave = registro
			cod_lib: N(6)
			cod_ejemplar: N(10)
		FReg
		descripcion: AN(255)
		tipo_libro: AN(1)
		disponible: ("SI","NO")
		estado: AN(10)
		fecha: Fecha 
	FReg
	arch1,sal: archivo de EJEMPLARES ordenado por clave
	reg1,reg_sal,aux: EJEMPLARES

	NOVEDADES = registro 
		clave = registro 
			cod_lib: N(6)
			cod_ejemplar: N(10)
		FReg
		tipo_novedad: (1...5)
		fecha_novedad: Fecha 
		id_usuario: N(20)
	FReg
	arch2: archivo de NOVEDADES ordenado por clave 
	reg2: NOVEDADES 

	USUARIOS = registro 
		id_usuario: N(20)
		nombre: AN(20)
		edad: N(2)
		tipo_usuario: ("E","D","P")
	FReg
	archInd: archivo de USUARIOS indexado por id_usuario 
	regInd: USUARIOS 

	suma,cont: entero 

	procedimiento num1()
		aux.cod_lib:= reg2.cod_lib
		aux.cod_ejemplar:= reg2.cod_ejemplar
		aux.disponible:= "SI"
		aux.fecha:= " "
		cont:= cont + 1
	FProcedimiento
	
	procedimiento num2()
		regInd.id_usuario:= reg2.id_usuario 
		Leer(archInd,regInd)
		si EXISTE entonces 
			aux.disponible:= "NO"
			segun regInd.tipo_usuario hacer 
				"E": suma:= 10 
				"D": suma:= 15 
				"P": suma:= 7
			fsegun 
			aux.fecha:= sumar_dias(reg2.fecha_novedad,suma)
		sino 
			Esc("Error usuario no existe")
		FSi 
	FProcedimiento

	procedimiento num3()
		aux.disponible:= "SI"
		aux.fecha:= " "
	FProcedimiento

	procedimiento num4()
		aux.disponible:= "NO"
		aux.estado:= "Reparacion"
		aux.fecha:= fecha_novedad
	FProcedimiento

	procedimiento num5()
		aux.disponible:= "NO"
		aux.estado:= "Baja"
		aux.fecha:= fecha_novedad
	FProcedimiento

	procedimiento procesos_iguales() 
		aux:= reg1
		Mientras aux.clave = reg2.clave hacer 
			segun reg2.tipo_novedad hacer 
				1: Esc("Error")
				2: num2()
				3: num3()
				4: num4()
				5: num5()
			FSegun
			Leer_nov() 
		FM
		reg_sal:= aux 
		Escribir(sal,reg_sal)
		Leer_ejem()
	FProcedimiento

	procedimiento procesos_distintos()
		si reg2.tipo_novedad = 1 entonces 
			num1()
		FSi 
		Leer_nov()
		Mientras aux.clave = reg2.clave hacer 
			procesos_iguales()
		FM 
		reg_sal:= aux 
		Escribir(sal,reg_sal)
	FProcedimiento

	procedimiento Leer_nov()
		Leer(arch1,reg1)
		si FDA(arch1) entonces 
			reg1.clave:= HV 
		FSi
	FProcedimiento

	procedimiento Leer_ejem()
		Leer(arch2,reg2)
		si FDA(arch2) entonces
			reg2.clave:= HV 
		FSi 
	FProcedimiento

	Proceso 
		Abrir E/ (arch1); Abrir E/(arch2); Abrir E/ (archInd); Abrir /S (sal)
		Leer_ejem(); Leer_nov()
		suma:= 0

		Mientras (reg1.clave <> HV) o (reg2.clave <> HV) hacer
			si reg1.clave < reg2.clave entonces 
				reg_sal:= reg1 
				Escribir(sal,reg_sal)
				Leer_ejem()
			sino 
				si reg1.clave = reg2.clave entonces 
					procesos_iguales()
				sino 
					si reg1.clave > reg2.clave entonces 
						procesos_distintos()
					FSi
				FSi
			FSi 
		FSi 
		Cerrar(arch1), Cerrar(arch2); Cerrar(archInd); Cerrar(sal)

		Esc("La cantidad de nuevas unidades es de: ",cont)
	FProceso 
FAccion 



