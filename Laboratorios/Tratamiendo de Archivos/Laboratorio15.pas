program lab2;
{Dado un archivo que almacena datos de restaurantes del mundo}

{Pregunta: Cuál es la cantidad de restaurantes que tienen calificación Excelente 
(Rating Text = Excellent), que ofrecen delivery online (Has Online delivery = Yes)}

{RESPUESTA: 39}

uses 
	csvdocument, sysutils; 
type 
	RESTAURANTES = record 
    	id_restaurante: Integer; 
    	restaurante: String[50]; 
    	cod_pais: String[30];
    	ciudad: String[30];
    	direccion: String[30];
    	localidad: String[30];
    	cocinas: String[30];
    	tiene_reserva: String[3];
    	delivery_online: String[3];
    	esta_entregando: String[3];
    	calificacion: Real;
    	color: String[15];
    	texto: String[15];
    	votos: Integer;
  	end; 
var
	arch: TCSVDocument; 
	reg: RESTAURANTES; 
	cant_cumplen, cant_total, i: Integer;
	porc: real;

	procedure inicializar();
	begin
		arch:= TCSVDocument.create();
		cant_cumplen:=0; cant_total:=0;
	end;
begin
	inicializar();
	arch.delimiter:= ';';
	arch.LoadFromFile('Pascal\TratamientoArchivos\Archivos\Laboratorio5.csv');

	for i:= 1 to (arch.rowcount - 1) do 
	begin
		reg.delivery_online:= arch.cells[8,i];
		reg.texto:= arch.cells[12,i];

		if (reg.delivery_online = 'Yes') and (reg.texto = 'Excellent') then 
		begin
			cant_cumplen+=1;
		end;
		cant_total+=1;
	end;

	arch.free;
	porc:= (cant_cumplen/cant_total)*100;
	writeln('La cantidad de restaurantes que tienen un rating Excelente y que ofrecen delivery online son en total: ',cant_cumplen);
	writeln('El porcentaje de restaurantes que cumplen con ambas condiciones sobre el total es: ',porc:0:2,'%');
	readln;
end.
