/*2. En un archivo secuencial Participantes se almacena datos de las olimpiadas (o algo asi):
PARTICIPANTES ordenado por Pais y Equipo.
Pais | Equipo | Atleta | tiempo (seg) | Disciplina ("100m" "400m" "4x100" "4x400") | Descalificacion ("si" "no")

se pide:
- Total del tiempo por Pais, Equipo y total general de los descalificados.
- Guardar en un archivo de salida RELEVOS el Pais y Equipo si el equipo sumó un tiempo menor
a 160 segundos en disciplinas ("4x100" "4x400"), tambien tienen que estar descalificados.*/

accion ejercicio2 es 
 Ambiente 
	PARTICIPANTES = registro 
		pais: AN(20)
		equipo: AN(20)
		atleta: AN(30)
		tiempo: N(4)
		disciplina: ("100m","400m","4x100","4x400")
		descalificado: ("si","no")
	FReg 

	RELEVOS = registro
		pais: AN(20)
		equipo: AN(20) 
	FReg

	arch: archivo de PARTICIPANTES ordenado por pais y equipo 
	reg: PARTICIPANTES
	arch_sal: archivo de RELEVOS
	reg_sal: RELEVOS 

	t_pais, t_equipo, t_general, t_4x100, t_4x400: entero 
	res_pais, res_equipo: AN(20)

	procedimiento corte_equipos() es 
		Esc("Total de tiempo del equipo ",res_equipo,": ",t_equipo)
		t_pais:= t_pais + t_equipo
		t_equipo:= 0 
		si (t_4x100 < 160) o (t_4x400 < 160) entonces 
			reg_sal.pais:= res_pais 
			reg_sal.equipo:= res_equipo
			Esc(arch_sal,reg_sal)
			t_4x100:= 0; t_4x400:= 0
		fsi 	
		res_equipo:= reg.equipo 
	FProcedimiento

	procedimiento corte_pais() es 
		corte_equipos()
		Esc("Total de tiempo del pais ",res_pais,": ",t_pais)
		t_general:= t_general + t_pais
		t_pais:= 0 
		res_pais:= reg.pais 
	FProcedimiento
	
	Proceso 
		Abrir E/(arch)
		Abrir S/(arch_sal)
		Leer(arch,reg)

		res_pais:= reg.pais; res_equipo:= reg.equipo 
		t_4x100:= 0; t_4x400:= 0; t_equipo:= 0; t_pais:= 0; t_general:= 0 

		Mientras NFDA(arch) hacer 
			si res_pais <> reg.pais entonces 
				corte_pais()
			sino 
				si res_equipo <> reg.equipo entonces 
					corte_equipo()
				fsi 
			fsi 

			si (reg.descalificado = "si") entonces 
				t_equipo:= t_equipo + reg.tiempo 
				si (reg.disciplina = "4x100") entonces 
					t_4x100:= t_4x100 + reg.tiempo
				sino 
					si (reg.disciplina = "4x400") entonces
						t_4x400:= t_4x400 + reg.tiempo
					fsi 
				fsi 
			fsi 
		FM 
		Cerrar(arch); Cerrar(arch_sal)

		Esc("El total general de tiempo es de: ",t_general)
	FProceso 
FAccion 