Accion ej1 es 
 Ambiente 
	NOVEDADES = registro 
		legajo: N(5)
		mes: N(2)
		año: N(4)
		cod_nov: AN(2)
		t_nov: ("D","R")
		monto: N(6,2)
	Freg
	arch_n: archivo de novedades ordenado por legajo, mes, año, cod_nov
	reg_n: NOVEDADES 

	EMPLEADOS = registro 
		legajo: N(5)
		nombre: AN(20)
		cargo: (1...10)
		basico: N(6,2)
	Freg 
	arch_e: archivo de EMPLEADOS indexado por legajo 
	reg_e: EMPLEADOS 

	SUELDOS = registro 
		legajo: N(5)
		mes: N(2)
		año: N(4)
		cargo: (1...10)
		neto: 
	Freg 
	arch_s: archivo de SUELDOS indexado por legajo, mes y año 
	reg_s,sal: SUELDOS 

	A: arreglo[1...10] de enteros 
	res_leg, cobrar: entero 
	
	Procedimiento actualizarCobrar()
		segun reg_n.t_nov hacer 
			"D": cobrar:= cobrar - reg_n.monto 
			"R": cobrar:= cobrar + reg_n.monto 
		Fsegun 
	FProcedimiento 

	Procedimiento initCobrar()
		reg_e.legajo:= res_leg
		Leer(arch_e,reg_e)
		si EXISTE entonces 
			cobrar:= reg_e.basico 
		Fsi 
	FProcedimiento

	Procedimiento actualizar()
		reg_e.legajo:= res_leg
		Leer(arch_e,reg_e) 
		si EXISTE entonces
			sal.legajo:= res_leg; sal.mes:= reg_e.mes; sal.nombre:= reg_e.nombre
			sal.año:= reg_e.año; sal.cargo:= reg_e.cargo; sal.neto:= cobrar 
			mostrarListado()
			A[reg_e.cargo]:= A[reg_e.cargo] + cobrar 
		Fsi 
	FProcedimiento

	Procedimiento mostrarListado()
		Esc("Legajo:",res_leg,"Nombre:",reg_e.nombre,"Cargo:",reg_e.cargo,"Mes:",reg_e.mes,"Año:",reg_e.año,"Neto:",cobrar)
	FProcedimiento

	Procedimiento init()
		Para i:= 1 hasta 10 hacer 
			A[i]:= 0 
		FPara 
		cobrar:= 0
		res_leg:= reg_n.legajo
	FProcedimiento

	Proceso 
		Abrir E/(arch_n); Leer(arch_n,reg_n)
		Abrir E/(arch_e); Abrir S/(arch_s)

		init()

		Mientras NFDA(arch_n) hacer 
			si cobrar = 0 entonces 
				initCobrar()
			sino 
				actualizarCobrar()
			Fsi 

			si res_leg <> reg_n.legajo entonces 
				actualizar()
				mostrarListado()
				cobrar:= 0
			sino 
				Leer(arch_n,reg_n)
			Fsi 
		FM 
		Cerrar(arch_n); Cerrar(arch_e); Cerrar(arch_s)

		Para i:= 1 hasta 10 hacer 
			Esc("Cargo",i,":",A[i],"$")
		FPara

	FProceso 
FAccion



