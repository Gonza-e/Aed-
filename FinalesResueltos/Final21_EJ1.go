accion ej1 es 
 Ambiente 
	Novedades = registro 
		legajo: N(5)
		mes: N(2)
		año: N(4)
		cod: AN(5)
		t_novedad: ("D","R")
		monto:
	Freg 
	arch_n: archivo de Novedades ordenado por legajo, mes, año, cod
	reg_n: Novedades

	Empleados = registro 
		legajo: N(5)
		nombre: AN(30)
		cargo: (1...10)
		basico: N(6,2)
	Freg
	arch_e: archivo de Empleados indexado por legajo 
	reg_e: Empleados 

	Sueldos = registro 
		legajo: N(5)
		mes: N(2)
		año: N(4)
		cargo: (1...10)
		neto: N(6,2)
	Freg 
	arch_s: archivo de Sueldos indexado por legajo, mes y año 
	reg_s,sal: Sueldos 

	Procedimiento entrarIndex()
		reg_e.legajo:= reg_n.legajo 
		Leer(arch_e,reg_e)
		si EXISTE entonces 
			cobrar:= reg_e.basico 
		Fsi 
	FProcedimiento 

	Procedimiento actSueldo()
		segun reg_n.t_novedad hacer 
			"D": cobrar:= cobrar - reg_n.monto 
			"R": cobrar:= cobrar + reg_n.monto 
		Fsegun 
	FProcedimiento

	Procedimiento listado()
		reg_s.legajo:= res_legajo
		Leer(arch_s,reg_s)
		si EXISTE entonces 
			Esc("Legajo:",res_legajo,"Nombre:",nom,"Cargo:",cargo,"Mes:",mes,"Año:",año,"Neto:",cobrar)
			sal.legajo:= res_legajo; sal.mes:= mes; sal.año:= año; sal.cargo:= cargo; reg_s.neto:= cobrar
			Esc(arch_s,sal)
			A[cargo]:= A[cargo] + cobrar 
			cobrar:= 0
		Fsi 
	FProcedimiento

	cobrar, res_legajo, cargo, mes, año: entero
	nom: AN 
	A: arreglo[1...10] de enteros
	
	Procedimiento init()
		Para i:= 1 hasta 10 hacer 
			A[i]:= 0 
		FPara 
	FProcedimiento

	Proceso 
		Abrir E/(arch_n); Leer(arch_n,reg_n)
		Abrir E/(arch_e); Abrir S/(arch_s)
		cobrar:= 0; res_legajo:= reg_n.legajo; nom:= ""
		init()

		Mientras NFDA(arch_n) hacer 
			si cobrar = 0 entonces 
				entrarIndex()
				nom:= reg_e.nombre; cargo:= reg_e.cargo; mes:= reg_n.mes; año:= reg_n.año
			sino 
				actSueldo()
			Fsi 

			si res_legajo <> reg_n.legajo entonces 
				listado()
			Fsi 

			Leer(arch_n,reg_n)
		FM 
		Cerrar(arch_n); Cerrar(arch_s); Cerrar(arch_e)

		Para i:= 1 hasta 10 hacer 
			Esc("Cargo:",i,"Monto:",A[i])
		FPara
	FProceso 
FAccion

