Program Laboratorio3; 

{* El Registro Civil posee datos de personas en el archivo "Laboratorio3.txt"
que contiene sus nombres (pueden ser varias palabras, indeterminado, finaliza
con "&") y fechas de nacimiento (ocho caracteres, ddmmaaaaa).

Se pide: ¿qué porcentaje de personas nacieron en la primera quincena (de cualquier
mes) en el primer trimestre (de cualquier año), cuyos nombres comienzan además
con letras consonantes? *}

uses 
 Funciones in '.../Funciones.pas';
var 
 sec: TextFile;
 v: Char; 
 total_personas, cumplen, dia, mes, i: Integer; 
 quincena, trimestre, consonante: Boolean; 
 porc: Real;

procedure inicializar();
begin
  total_personas:= 0; cumplen:= 0; dia:= 0; mes:= 0; porc:= 0; 
  quincena:= False; trimestre:= False; consonante:= False;
end;

begin
  Assign(sec,'ArchivosParaUso/Laboratorio3.txt');
  Reset(sec);
  Read(sec,v);
  inicializar; 
  while not Eof(sec) do
  begin

    //WriteLn(v);
    if esConsonante(v) then
    begin
      consonante:= True;
    end;

    while (v <> '&') do
    begin
      Read(sec,v);
    end;

    Read(sec,v);
    dia:= conv_entero(v) * 10; 
    Read(sec,v); 
    dia:= dia + conv_entero(v); 
    //WriteLn(dia);

    if dia <= 15 then 
    begin
      quincena:= True;
    end;

    Read(sec,v);
    mes:= conv_entero(v) * 10;
    Read(sec,v);
    mes:= mes + conv_entero(v); 
    //WriteLn(mes);

    if mes <= 3 then 
    begin
      trimestre:= True;
    end;

    for i:= 1 to 5 do 
    begin
      Read(sec,v);
    end;

    if consonante and trimestre and quincena then 
    begin
      cumplen:= cumplen + 1;
    end;
    total_personas:= total_personas + 1;
    consonante:= False; quincena:= False; trimestre:= False;
  end;

  porc:= (cumplen/total_personas) * 100;
  WriteLn('El porcentaje de personas que cumplen con los requisitos es de: ',porc:0:2,'%');
  WriteLn('Personas que cumplen: ',cumplen);
  WriteLn('Personas totales: ',total_personas);
  Close(sec);
end.
