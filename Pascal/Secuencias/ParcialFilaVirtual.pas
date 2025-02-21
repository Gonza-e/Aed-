Program ParcialFilaVirtual; 

uses 
 Funciones in '.../Funciones.pas';
var 
 sec1, sec2, salida: TextFile;
 v1, v2: Char;
 cant_abandono, i, hora: Integer; 
 bandera, bandera2, bandera3: Boolean;
begin
  Assign(sec1,'ArchivosParaUso/fila_virtual.txt');
  Reset(sec1);
  Assign(sec2,'ArchivosParaUso/compras.txt');
  Reset(sec2);
  Assign(salida,'secuenciaSalida.txt');
  Rewrite(salida); 
  cant_abandono:= 0; hora:= 0;
  bandera:= False; bandera2:= False; bandera3:= False;
  while (not Eof(sec1)) and (not Eof(sec2)) do
  begin
    Read(sec2,v2);
    Read(sec1,v1);
    hora:= conv_entero(v1) * 10; 
    Read(sec1,v1);
    hora:= hora + conv_entero(v1);
    while (not Eof(sec1)) and (v1 <> '#') do
    begin
      if hora < 10 then
      begin
        bandera3:= True
      end;
        while (not Eof(sec2)) and (v2 <> '?') do
        begin
          while (not Eof(sec2)) and (v2 <> '+') and (v2 <> '.') and (v2 <> '?') do
          begin
            if letras(v2) and bandera3 then
            begin
              Write(salida,v2);
              bandera2:= True;
              bandera:= True;
            end
            else if v2 = '#' then
            begin
              cant_abandono:= cant_abandono + 1
            end;
            Read(sec2,v2);   
          end;
          Read(sec2,v2);
          if bandera then
          begin
            Write(salida,'+');
            bandera:= False;
          end;
        end;
        if bandera2 then
        begin
          Write(salida,'#');
          bandera2:= False;
          bandera3:= False;
        end;
        Read(sec1,v1);   
    end;
  end;
  Close(sec1);
  Close(sec2);
  Close(salida);
  WriteLn('La cantidad de usuarios que abandonaron es de: ',cant_abandono);
end.