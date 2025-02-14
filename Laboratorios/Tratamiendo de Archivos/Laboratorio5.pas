Program Laboratorio5; 

{* Se tiene un archivo (entrada-lab-ej5.csv) que contiene datos de restaurantes
del mundo con la siguiente información:

(*) ID del restaurante: N(12).
(*) Nombre del restaurante: AN(200).
(*) Código de país: N(3).
(*) Ciudad: AN(200).
(*) Dirección: AN(200).
(*) Localidad: AN(200).
(*) Cocina: AN(200).
(*) Tiene reserva de mesa: ("YES", "NO").
(*) Tiene delivery online: ("YES", "NO").
(*) Está realizando entregas: ("YES", "NO").
(*) Calificación agregada: N(4,2).
(*) Color de la calificación: AN(30).
(*) Texto de la calificación: ("Excellent", "Very Good", "Good", "Poor",
    "Average", "Not rated").
(*) Votos: N(3).

Se pide: ¿cuál es la cantidad de restaurantes que no tienen reserva de mesa (Has
table booking = No) pero sí ofrecen delivery online (Has online delivery = Yes)? *}

uses 
 csvdocument, 
 sysutils,
 Funciones in 'Pascal/Secuencias/Funciones.pas';
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
 sicumple, i: Integer;
 procedure inicializar();
 begin
   sicumple:= 0; 
   arch:= TCSVDocument.create();
 end;
begin
  inicializar(); 
  arch.delimiter:= ';';
  arch.LoadFromFile('C:\Users\Gonzalo\Desktop\Algoritmos\2024\Ejercicios de la guia\Pascal\TratamientoArchivos\Archivos\laboratorio5.csv');

  
  for i:= 1 to (arch.rowcount - 1) do 
  begin
    if arch.cells[7,i] <> ' ' then 
    begin
      reg.tiene_reserva:= arch.cells[7,i];
    end;
   { else 
    begin
      reg.tiene_reserva:= ' ';
    end;}

    if arch.cells[8,i] <> ' ' then 
    begin
      reg.delivery_online:= arch.cells[8,i]; 
    end;
   { else 
    begin
      reg.delivery_online:= ' ';
    end;}

    if (reg.tiene_reserva = 'No') and (reg.delivery_online = 'Yes') then 
    begin
      //sicumple:= sicumple + 1;
      //inc(sicumple);
      sicumple+=1;
    end;
  end;
  WriteLn('La cantidad de restaurantes que cumplen con las condiciones es de: ',sicumple);
  arch.free;
  ReadLn;
end.
