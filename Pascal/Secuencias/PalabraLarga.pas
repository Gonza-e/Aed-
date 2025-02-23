Program PalabraLarga; 

uses 
 Funciones in '.../Funciones.pas';
var
 sec: TextFile;
 v: Char;
 palabra_larga, cont_letras: Integer; 
 bandera: Boolean;

begin
  Assign(sec,'ArchivosParaUso/PROBLEMA3.txt');
  Reset(sec);
  palabra_larga:= 0; cont_letras:= 0;
  bandera:= False; 
  v:= ',';
  while not Eof(sec) do
  begin
    if enPalabra(v) then 
    begin
      bandera:= True;
      cont_letras:= cont_letras + 1; 
    end
    else if bandera then 
    begin
      bandera:= False; 
      if cont_letras > palabra_larga then
      begin
        palabra_larga:= cont_letras;
      end;
      cont_letras:= 0;
    end;
    Read(sec,v);   
  end;
  Close(sec);
  WriteLn('La palabra mas larga cuenta con: ',palabra_larga,' letras');

  readln;
end.
