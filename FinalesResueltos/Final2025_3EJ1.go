accion ejercicio1 es 
 Ambiente 
	sec: secuencia de caracteres 
	v: caracter

	nodo = registro 
		id: N(5)
		titulo: AN(40)
		area: N(2)
		tipo: N(1)
		estado: AN(1)
		prox: puntero a nodo 
	Freg
	prim,p,q,ant: puntero a nodo 

	area_may, may_proy, res_area, res_tipo, cant_proy: entero 
	res_id, res_titulo, res_estado: AN 

	AREAS = registro 
		cod_area: N(2)
		nom_area: AN(40)
	Freg
	arch: archivo de AREAS indexado por cod_area
	reg: AREAS 

	procedimiento cargaOrdenada() es 
		p:= prim; ant:= nil 
		Mientras (p <> nil) y (*q.tipo > *p.tipo) o  ((*q.tipo = *p.tipo) Y (*q.area > *p.area)) hacer 
			ant:= p 
			p:= *p.prox 
		FM 

		si p = prim entonces 
			*q.prox:= p 
			prim:= q 
		sino 
			*ant.prox:= q
			*q.prox:= p 
		fsi 
	FProcedimiento

	procedimiento obtenerArea() es 
		reg.cod_area:= *p.area
		Leer(arch,reg)
		si EXISTE entonces 
			Esc(reg.nom_area)
		fsi  
	FProcedimiento

	Proceso 
		ARR(sec); AVZ(sec,v)
		Abrir E/(arch)
		area_may:= 0; may_proy:= 0 
		prim:= nil

		Mientras NFDS(sec) hacer 
			Para i:= 1 a 5 hacer 
				Concatenar(res_id,v)
				AVZ(sec,v)
			FPara 

			Mientras v <> "&" hacer 
				Concatenar(res_titulo,v)
				AVZ(sec,v)
			FM 
			AVZ(sec,v)

			Para i:= 1 a 2 hacer 
				res_area:= (res_area * 10) + conv_entero(v)
				AVZ(sec,v)
			FPara

			res_tipo:= conv_entero(v)
			AVZ(sec,v)

			res_estado:= v 
			AVZ(sec,v) 

			Nuevo(q); *q.id:= res_id; *q.titulo:= res_titulo; *q.area:= res_area; *q.tipo:= res_tipo; *q.estado:= res_estado
			cargaOrdenada()
			res_id:= ' '; res_titulo:= ' '; res_area:= 0; res_tipo:= 0; res_estado:= ' '
		FM 

		p:= prim; s:= p 
		Mientras (p <> nil) hacer 
			Mientras *p.area = *s.area hacer 
				cant_proy:= cant_proy + 1 
				s:= *s.prox 
				Esc("Area: ",obtenerArea(*s.area),"Tipo: ",*s.area,"Titulo: ",*s.titulo)
			FM 

			si cant_proy > may_proy entonces 
				may_proy:= cant_proy
				cant_proy:= 0 
				area_may:= *p.area 
			fsi 
			p:= s 
		FM 

		Cerrar(sec)
		Esc("Area con mayor cantidad de proyectos: ",area_may) // consigna 1
	FProceso 
FAccion 
