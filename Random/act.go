Accion ej1 es 
	Ambiente 

	PILOTOS = registro 
		nom: AN(30)
		equipo: AN(30)
		p_acum: N(3)
		c_vrapidas: N(1)
		max_v: N(3)
	Freg 
	arch_p,sal: archivo de PILOTOS ordenado por nom 
	reg_p,reg_sal,aux: PILOTOS

	CARRERAS = registro 
		nom: AN(30)
		id_carrera: N(2)
		equipo: AN(30)
		ciudad: AN(20)
		pos: (1...20)
		v_rapida: booleano
		max_v: N(3)
		abandono: ("S","N")
	Freg 
	arch_c: archivo de CARRERAS ordenado por nom e id_carrera
	reg_c: CARRERAS

	ABANDONOS = registro 
		nom: AN(30)
		id_carrera: N(2)
		razon: AN(255)
		vuelta: N(2)
	Freg 
	arch: archivo de ABANDONOS indexado por nom e id_carrera
	reg: ABANDONOS

	resNom,equipo: AN 
	max,vmax,puntos,vueltas: entero

	nodo = registro 
		id: N(2)
		total: N(2)
		abandono: N(2)
		prox: puntero a nodo 
	Freg 
	prim,p,q: puntero a nodo

	Procedimiento act()
		Mientras resNom = reg_c.nom hacer
			si reg_c.abandono = "S" entonces 
				reg.nom:= resNom; reg.id_carrera:= reg_c.id_carrera
				Leer(arch,reg)
				si EXISTE entonces 
					Esc("Abandono: ",reg.nom," Ciudad: ",reg_c.ciudad,"Razon: ",reg.razon)
				Fsi 
			sino 
				si max < reg_c.max_v entonces 
					max:= reg_c.max_v
				Fsi 
			
				puntos:= puntos_obtenidos(reg_c.pos) + puntos 

				si reg_c.v_rapida y reg_c.pos <= 10 entonces 
					vueltas:= vueltas+1 
				Fsi 

				si reg_c.pos = 1 entonces 
					Esc("Ganador: ",reg_c.nom)
				Fsi 
			Fsi 
			actLista()
			Leer(arch_c,reg_c)
		FM 

		si max > vmax entonces 
			vmax:= max 
		Fsi 

		puntos:= puntos + vueltas 

	FProcedimiento 

	Procedimiento iguales()
		aux:= reg_p 
		act()
		aux.p_acum:= puntos; aux.c_vrapidas:= vueltas; aux.max_v:= max 
	FProcedimiento

	Procedimiento distintos()
		resNom:= reg_c.nom; equipo:= reg_c.equipo 
		act()
		aux.nom:= resNom; aux.equipo:= equipo; aux.p_acum:= puntos; aux.c_vrapidas:= vueltas; aux.max_v:= max 
	FProcedimiento

	Procedimiento actLista()
		si prim = nil entonces 
			Nuevo(q); *q.id:= reg_c.id_carrera; *q.total:= 1
			si reg_c.abandono = "S" entonces 
				*q.abandono:= 1 
			Fsi 
		sino 
			p:= prim 
			Mientras p<>nil y *q.id <> *p.id hacer 
				p:= *p.prox 
			FM 

			si p = nil entonces 
				Nuevo(q); *q.id:= reg_c.id_carrera; *q.total:= 1 
				si reg_c.abandono = "S" entonces 
					*q.abandono:= 1 
				Fsi 
			sino 
				*q.total:= *q.total+1; *q.abandono:= *q.abandono+1 
			Fsi 
		Fsi 
	FProcedimiento

	Proceso 
		Abrir E/(arch_p); Leer(arch_p,reg_p)
		Abrir S/(sal); Abrir E/(arch)

		puntos:= 0; vueltas:= 0; max:= 0; vmax:= 0
		resNom:= reg_p.nom; equipo:= " "

		Mientras NFDA(arch_p) y NFDA(arch_c) hacer 
			si resNom <> reg_c.nom entonces 
				distintos()
			sino 
				iguales()
			Fsi 
			reg_sal:= aux 
			Grabar(sal,reg_sal)
			puntos:= 0; vueltas:= 0; max:= 0

			Leer(arch_p,reg_p)
		FM 
		Cerrar(arch); Cerrar(arch_p); Cerrar(arch_c)

		p:= prim 
		Mientras p<>nil hacer 
			Esc("Carrera: ",*p.id,"Promedio de abandono: ",((*p.abandono)/(*p.total))*100)
			p:= *p.prox 
		FM 
		
		Esc("La velocidad mas alta registrada fue: ",vmax)

	FProceso 
FAccion