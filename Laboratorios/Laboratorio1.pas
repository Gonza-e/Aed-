Program Laboratorio1; 

uses 

  Funciones in '.../Funciones.pas',
  math;  //Coloco math para poder usar potencias 

var 

  sec: TextFile; 
  v: Char; 
  cant_supera, i, k: Integer; 
  importe: LongInt;

begin
  Assign(sec,'ArchivosParaUso/Laboratorio1.txt');
  Reset(sec);
  Read(sec,v); 
  cant_supera:= 0; k:= 3; importe:= 0;
  while not Eof(sec) do
  begin
    For i:= 1 to 4 do 
    begin
      Read(sec,v);
    end;

    For i:= 1 to 3 do 
    begin
      importe:= importe + conv_entero(v) * (10 ** (k - 1));
      Read(sec,v);
    end;

    if importe > 400 then
    begin
      cant_supera:= cant_supera + 1;
    end;
    importe:= 0;

  end;
  WriteLn('La cantidad de productos que superan los 400$ es de: ',cant_supera);
  Close(sec);
end.