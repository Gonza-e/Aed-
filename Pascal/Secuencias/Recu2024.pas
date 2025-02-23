Program Recu2024; 

uses
 Funciones in '.../Funciones.pas';
var
  sec1, sec2, sal: TextFile;
  v1, v2: char;
  cant_participantes, i, sisupera, nosupera, edad, edad_comp: Integer;
begin
  Assign(sec1,'ArchivosParaUso/pais.txt');
  Assign(sec2,'ArchivosParaUso/secuencia.txt');
  Reset(sec1); Reset(sec2);
  Assign(sal,'secuencia de salida');
  Rewrite(sal);
  Read(sec1,v1); Read(sec2,v2);
  cant_participantes:= 0; sisupera:= 0; nosupera:= 0; edad:= 0; edad_comp:= 0;

  WriteLn('Ingrese edad para comparar: ');
  Read(edad);

  while (not Eof(sec1)) and (not Eof(sec2)) do
  begin
    while (v1 <> '@') do
    begin
      if v1 = 'P' then
      begin
        Read(sec1,v1);
        while (v1 <> '#') do
        begin
          Write(sal, v1);
          Read(sec1, v1);
        end;
        Write(sal,'#');
        while (not Eof(sec2)) and (v2 <> '@') do
        begin
          for i:= 1 to 4 do 
          begin
            Read(sec2,v2);
          end;
          edad_comp:= conv_entero(v2) * 10;
          Read(sec2,v2);
          edad_comp:= edad_comp + conv_entero(v2);
          Read(sec2,v2);

          if edad > edad_comp then
          begin
            nosupera:= nosupera + 1;
          end
          else
          begin
            sisupera:= sisupera + 1;
          end;

          while (not Eof(sec2)) and (v2 <> '#') and (v2 <> '@') do
          begin
            Write(sal,v2);
            Read(sec2,v2);
          end;
          cant_participantes:= cant_participantes + 1;
          Write(sal,'%');
        end;
      Read(sec2,v2);
      end; 
      Read(sec1,v1); 
    end;
  end;
  Close(sec1);
  Close(sec2);
  Close(sal);
  WriteLn('La cantidad de participantes que superan la edad es de: ',sisupera);
  WriteLn('La cantidad de participantes que no superan la edad es de: ',nosupera);
  WriteLn('La cantidad de participantes de estas delegaciones son: ',cant_participantes);

  readln;
end.
