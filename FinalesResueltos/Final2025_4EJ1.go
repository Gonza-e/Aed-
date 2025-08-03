/*1. Se tiene almacenado los datos de pacientes en una secuencia de caracteres
Historia clinica(4 caracteres)Nombre-TipoDeCobertura(de 10 a 29)TipoDeEspecialidad(0 a 9)Dia('L','M','V')

diseñar un algoritmo que muestre un menu con la siguientes acciones :
    - separar los datos en HistoriaClinica y Tipo de Especialidad  (no me acuerdo si eran exactamente esos dato, pero si pedia 2).
    - buscar a un paciente y si se encuentra, informar el Historia Clinica, tipo de Cobertura y Dia.
    - informar: 
        - total por tipo de especialidad y tipo de cobertura
        - el tipo de especialidad que mas se solicitó
        - el tipo de cobertura que menos se registró.
    Se puede usar la funcion CONCATENAR ----> la funcion concantena caracteres, los parametros son la variable donde guardar
	y el caracter a concatenar (ejemplo:= CONCATENAR (ejemplo,v)). Se usa dentro de un ciclo.

Utilizar la estructura que crea conveniente*/

Accion ejercicio1 es 
 Ambiente 
    sec: secuencia de caracteres
    v: caracter
    nombre, hist, d: AN 
    t_cobertura, t_especialidad, pos_may, e_may, pos_men, c_men: entero 

    nodo = registro 
        nombre: AN(30)
        hist_c: AN(4)
        tipo_cobertura: N(2)
        tipo_especialidad: N(1)
        dia: AN(1)
        prox: puntero a nodo 
    FReg 
    p,q,ant,prim: puntero a nodo 
    
    A: arreglo[1...19] de enteros 
    B: arreglo[1...10] de enteros 

    funcion conv_entero(v:caracter):entero 
        ="0": conv_entero:= 0
        ="1": conv_entero:= 1
        ="2": conv_entero:= 2
        ="3": conv_entero:= 3
        ="4": conv_entero:= 4
        ="5": conv_entero:= 5
        ="6": conv_entero:= 6
        ="7": conv_entero:= 7
        ="8": conv_entero:= 8
        ="9": conv_entero:= 9
    FFuncion

    procedimiento carga_ordenada() es 
        p:= prim; ant:= nil 
        Mientras (p<>nil) y ((*q.hist_c > *p.hist_c) y (*q.t_especialidad > *p.t_especialidad)) hacer 
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

    Proceso 
        ARR(sec); AVZ(sec)
        prim:= nil 
        Para i:= 1 a 19 hacer 
            A[i]:= 0 
        FPara 
        Para i:= 1 a 10 hacer 
            B[i]:= 0 
        FPara 
        nombre:= " "; t_cobertura:= 0; hist:= " "; t_especialidad:= 0; d:= " "
        pos_may:= 0; pos_men:= 0; e_may:= LV; c_men:= HV 

        Mientras NFDS(sec) hacer 
            Para i:= 1 a 4 hacer 
                concatenar(hist,v)
                AVZ(sec,v)
            FPara

            Mientras v <> "-" hacer 
                concatenar(nombre,v)
                AVZ(sec,v)
            FM
            AVZ(sec,v)
            
            Para i:= 1 a 2 hacer 
                t_cobertura:= (t_cobertura * 10) + conv_entero(v)
                AVZ(sec,v)
            FPara 

            t_especialidad:= conv_entero(v) 
            AVZ(sec,v)
            d:= v 
            AVZ(sec,v)

            A[t_cobertura]:= A[t_cobertura] + 1 
            B[t_especialidad]:= B[t_especialidad] + 1  

            nuevo(q); *q.nombre:= nombre; *q.hist_c:= hist; *q.tipo_cobertura:= t_cobertura; *q.dia:= d 
            si prim = nil entonces
                *q.prox:= prim 
                prim:= q 
            sino 
                carga_ordenada()
            fsi 

            t_cobertura:= 0
        FM 

        Cerrar(sec)
        Esc("Ingrese el nombre del paciente")
        Leer(nom_paciente)

        p:= prim
        Mientras p <> nil y (nom_paciente <> *p.nombre) hacer
            p:= *p.prox 
        FM 
        si p = nil entonces 
            Esc("El paciente no se encontro")
        sino 
            Esc("Historia clinica: ",*p.hist," Tipo de cobertura: ",*p.tipo_cobertura," Dia: ",*p.dia)
        fsi 
         
        Para i:= 1 a 19 hacer 
            Esc("Especialidad: ",i," Cantidad: ",A[i])
            si e_may < A[i] entonces 
                e_may:= A[i]
                pos_may:= i
            fsi 
        FPara 
        Para i:= 1 a 10 hacer 
            Esc("Cobertura: ",i," Cantidad: ",B[i])
            si c_men > B[i] entonces 
                c_men:= B[i]
                pos_men:= i
            fsi 
        FPara 

        Esc("La especialidad mayor fue: ",pos_may); Esc("La cobertura menor fue: ",pos_men)
    FProceso 
FAccion 



