Program CantidadPalabras; 

uses
 Funciones in '.../Funciones.pas';
var 
 sec: TextFile;
 v: Char; 
 cant_palabras: Integer; 
begin
  Assign(sec,'ArchivosParaUso/PROBLEMA3.txt');
  Reset(sec); 
  cant_palabras:= 0;
  v:= ' '; 
  while not Eof(sec) do
  begin
    if (v = ' ') or (v = '.') then
    begin
      cant_palabras:= cant_palabras + 1;
      Read(sec,v);
    end;
    Read(sec,v);
  end;
  Close(sec);
  WriteLn('La cantidad de palabras del texto es de: ',cant_palabras);
end.