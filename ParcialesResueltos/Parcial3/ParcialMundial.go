Accion mundial (prim: puntero a nodo) es 
 Ambiente 
 	nodo = registro
	 pais: AN(20)
	 grupo: AN(1)
	 puntos: N(2)
	 dif_gol: N(2)
	 t_amarillas: N(2)
	 t_rojas: N(2)
	 prox: puntero a nodo 
	FReg
 p, e_local, e_visitante: puntero a nodo 

 	nodo1 = registro
	 pais: AN(20)
	 grupo: AN(1)
	 puntos: N(2)
	 dif_gol: N(2)
	 t_amarillas: N(2)
	 t_rojas: N(2)
	 prox, ant: puntero a nodo1 
	FReg
 prim1, p1, q1: puntero a nodo1 

	PARTIDOS = registro 
	 id_partido: AN(10)
	 equipo_local: AN(20)
	 equipo_visit: AN(20)
	FReg
 arch: archivo de PARTIDOS 
 reg: PARTIDOS
	
	RESULTADOS = registro 
	 id_partido: AN(10)
	 cant_goles_local: N(2)
	 cant_goles_visit: N(2)
	 t_amarillas: N(2)
	 t_rojas; N(2)
	FReg
 arch_ix: archivo de RESULTADOS indexado por id_partido
 reg_ix: RESULTADOS

	Procedimiento actualizar() es 
	 	Si reg_ix.cant_goles_local < reg_ix.cant_goles_visit entonces 
		 *e_local.dif_gol:= *e_local.dif_gol + (reg_ix.cant_goles_local - reg_ix.cant_goles_visit)
		 *e_visitante.dif_gol:= *e_visitante.dif_gol + (reg_ix.cant_goles_local - reg_ix.cant_goles_visit)
		 *e_visitante.puntos:= *e_visitante.puntos + 3 
		Sino 
			Si reg_ix.cant_goles_local = reg_ix.cant_goles_visit entonces 
			 *e_local.puntos:= *e_local.puntos + 1 
			 *e_visitante:= *e_visitante + 1 
			Sino 
			 *e_local.dif_gol:= *e_local.dif_gol + (reg_ix.cant_goles_local - reg_ix.cant_goles_visit)
			 *e_visitante.dif_gol:= *e_visitante.dif_gol + (reg_ix.cant_goles_local - reg_ix.cant_goles_visit)
			 *e_local.puntos:= *e_local.puntos + 1 
			FSi 
		FSi
	FProcedimiento

	Procedimiento buscar_equipo(nombre: AN(20), x: puntero a nodo) es 
	 p:= prim 
	 Mientras (p<>nil) Y (*p.pais <> nombre) hacer 
	 	 p:= *p.prox 
		FM 
	 x:= p 
	FProcedimiento

 Proceso 
	 Abrir E/(arch)
	 Abrir E/(arch_ix)
	 Leer(arch,reg)
	 prim1:= nil
	 Mientras NFDA(arch) hacer 
	 	 reg_ix.id_partido:= reg.id_partido 
		 Leer(arch_ix,reg_ix)
		 	Si EXISTE entonces 
			 buscar_equipo(reg.equipo_local, e_local)
			 buscar_equipo(reg.equipo_visit, e_visitante)
			 actualizar()
			FSi 
		 Leer(arch,reg)
		FM 
	 p:= prim 
	 Mientras (p<>nil) hacer 
	 	 	Si (*p.puntos > 4) O ((*p.puntos = 4) Y (*p.dif_gol > 2) O ) entonces 
			 cargar_lista_doble()
			FSi 
		 p:= *p.prox 
		FM 

					
		 	
