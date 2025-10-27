accion 2025 es 
 ambiente 
	LIBROS = registro 
		cod_lib: N(8)
		titulo: AN(60)
		ISBN: N(13)
		descripcion: AN(255)
		tipo_libro: AN(1)
		cant_dias_max: N(2)
	FReg
	archInd: archivo de LIBROS indexado por cod_lib 
	regInd: LIBROS 

	EJEMPLARES = registro 
		clave = registro
			cod_lib: N(6)
			cod_ejemplar: N(10)
		FReg
		disponible: ("SI","NO")
		estado: AN(10)
		fecha: Fecha 
	FReg
	arch1,arch_sal: archivo de EJEMPLARES ordenado por clave 
	reg1,aux,sal: EJEMPLARES

	NOVEDADES = registro 
		clave = registro 
			cod_lib: N(6)
			cod_ejemplar: N(10)
			tipo_novedad: (1...5)
			fecha_novedad: Fecha 
		FReg
		id_usuario: N(20)
	FReg
	arch2: archivo de NOVEDADES ordenado por clave 
	reg2: NOVEDADES 

	cont: entero 

	procedimiento num1() 
		aux.cod_lib:= reg2.cod_lib 
		aux.cod_ejemplar:= reg2.cod_ejemplar
		aux.disponible:= "SI"
		aux.fecha:= " "
		cont:= cont + 1
	FProcedimiento

	procedimiento num2()
		aux.disponible:= "NO"
		regInd.cod_lib:= reg2.cod_lib
		Leer(archInd,regInd)
		si EXISTE entonces 
			aux.fecha:= sumar_dias(reg2.fecha_novedad,regInd.cant_dias_max)
		sino 
			Esc("Error, no existe el libro")
		FSi
	FProcedimiento

	procedimiento num3()
		aux.disponible:= "SI"
		aux.fecha:= " "
	FProcedimiento

	procedimiento num4()
		aux.disponible:= "NO"
		aux.estado:= "Reparacion"
		aux.fecha:= reg2.fecha_novedad
	FProcedimiento

	procedimiento num5()
		aux.disponible:= "NO"
		aux.estado:= "Baja"
		aux.fecha:= reg2.fecha_novedad
	FProcedimiento

	procedimiento Leer_ejem()
		Leer(arch1,reg1)
		si FDA(arch1) entonces 
			reg1.clave:= HV 
		FSi
	FProcedimiento

	procedimiento Leer_nov()
		Leer(arch2,reg2)
		si FDA(arch2) entonces 
			reg2.clave:= HV 
		FSi 
	FProcedimiento

	procedimiento procesos_iguales() 
		aux:= reg1 
		Mientras aux.clave = reg2.clave hacer 
			segun reg2.tipo_novedad hacer 
				1: Esc("Error")
				2: num2()
				3: num3()
				4: num4():
				5: num5():
			FSegun 
			Leer_nov()
		FM 
		sal:= aux 
		Escribir(arch_sal,sal)
		Leer_ejem()
	FProcedimiento

	procedimiento procesos_distintos()
		si reg2.tipo_novedad = 1 entonces 
			num1()
		FSi
		Leer_nov()
		procesos_iguales()
	FProcedimiento

	Proceso 
		Abrir E/ (arch1); Abrir E/ (arch2); Abrir E/(arch_sal)
		Leer_ejem(); Leer_nov()
		cont:= 0 

		Mientras (reg1.clave <> HV) o (reg2.clave <> HV) hacer
			si reg1.clave < reg2.clave entonces 
				sal:= reg1 
				Escribir(arch_sal,sal)
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
		FM 
		Cerrar(arch1); Cerrar(arch2); Cerrar(arch_sal)

		Esc("La cantidad de ejemplares nuevos es de: ",cont)
	FProceso
FAccion

Accion Ejercicio2 es 
 Ambiente 
	Fecha = registro 
		dia: N(2)
		mes: N(2)
		año: N(4)
	FReg

	NOVEDADES = registro 
		clave = registro 
			cod_lib: N(6)
			cod_ejemplar: N(10)
			tipo_novedad: (1...5)
			fecha_novedad: Fecha 
		FReg
		id_usuario: N(20)
	FReg
	arch1: archivo de NOVEDADES ordenado por clave 
	reg1: NOVEDADES 

	Funcion obtenerNovedad(x:entero):AN 
		segun x hacer 
			1: obtenerNovedad:= "Ingreso"
			2: obtenerNovedad:= "Prestamo"
			3: obtenerNovedad:= "Devolucion"
			4: obtenerNovedad:= "Mantenimiento"
			4: obtenerNovedad:= "Baja"
		FSegun
	FFuncion 

	Funcion obtenerMes(t:entero):AN 
		segun t hacer 
			1: obtenerMes:= "Enero"
			2: obtenerMes:= "Febrero"
			3: obtenerMes:= "Marzo"
			4: obtenerMes:= "Abril"
			5: obtenerMes:= "Mayo"
			6: obtenerMes:= "Junio"
		FSegun
	FProcedimiento

	A: arreglo[1...6,1...7] de enteros 
	mesMay, prestMay, i, j: entero 

	Proceso 
		Abrir E/ (arch1); Leer(arch1,reg1)

		Para i:= 1 hasta 6 hacer 
			Para j:= 1 hasta 7 hacer 
				A[i,j]:= 0
			FPara
		FPara
		
		mesMay:= 0; prestMay:= LV 

		Mientras NFDA(arch1) hacer 
			si (reg1.fecha_novedad.mes = 1) o (reg1.fecha_novedad.mes = 6) entonces 
				i:= reg1.tipo_novedad
				j:= reg1.fecha_novedad.mes 	

				A[i,j]:= A[i,j] + 1 
				A[6,j]:= A[6,j] + 1 
				A[i,7]:= A[i,7] + 1 
				A[6,7]:= A[6,7] + 1 

			Fsi 
			Leer(arch1,reg1)
		FM 

		Para i:= 1 hasta 6 hacer 
			Para j:= 1 hasta 7 hacer 
				si (i <> 6) y (j <> 7) entonces 
					si (prestMay < A[i,j]) y (i = 2) entonces 
						prestMay:= A[i,j]
						mesMay:= j 
					FSi
				FSi
				
				si i = 6 entonces 
					Esc("Totales por novedad: ",A[i,j])
				FSi 
				si j = 7 entonces 
					Esc("Totales por mes: ",A[i,j])
				FSi 
			FPara 
		FPara 
		Cerrar(arch1,reg1)

		Esc("El mes de ",obtenerMes(mesMay)," se tuvo una cantidad de prestamos de: ",prestMay)
	FProceso 
FAccion