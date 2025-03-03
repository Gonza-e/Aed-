Program ParcialTema7; 

uses 
  Funciones in '../Funciones.pas';
var
  sec: TextFile;
  v: char;
  sal: TextFile;
  cant_estudios, cant_general, cant_pacientes, recaudado, i, k: Integer;
  prom: Real;
  bandera: Boolean;

  begin
    Assign(sec,'ArchivosParaUso/secuenciaFarmacia.txt');
    Reset(sec);
    Assign(sal,'secuenciaSalida');
    Rewrite(sal);
    v:= ' '; 
    k:= 4;
    bandera:= False;
    cant_estudios:= 0; prom:= 0; cant_general:= 0; cant_pacientes:= 0; recaudado:= 0;
    while not Eof(sec) do
    begin
      while (not Eof(sec)) and (v <> '*') do
      begin
        while (not Eof(sec)) and (v <> '*') and (v <> ',') do
        begin
        //  Write(sal,v);
          Read(sec,v);
        end;
      cant_pacientes:= cant_pacientes + 1;
      Read(sec,v);
      cant_estudios:= conv_entero(v) * 10; 
      Read(sec,v);
      cant_estudios:= cant_estudios + conv_entero(v);
      Read(sec,v);
      WriteLn(cant_estudios);
      cant_general:= cant_general + cant_estudios;
      if cant_estudios >= 2 then
      begin
        For i:= 1 to (cant_estudios * 4) do 
        begin
          if v = 'A' then 
          begin
            bandera:= True;
            recaudado:= recaudado + 300;
          end
          else if v = 'E' then 
          begin
            recaudado:= recaudado + 420;
          end
          else if v = 'I' then 
          begin
            recaudado:= recaudado + 670;
          end;
          if bandera then
          begin
            k:= k-1;
            Write(sal,v);
           // Read(sec,v);
          end;
          if k = 0 then
          begin
            k:= 4;
            Write(sal,' / ');
            bandera:= False;
          end;
          Read(sec,v);
        end;
      end; 
      end;
    end;
    Close(sec);
    Close(sal);
    prom:= cant_general / cant_pacientes;
    WriteLn('El total recaudado fue de: ',recaudado,'$');
    WriteLn('El promedio de estudios por paciente es de: ',prom:0:1);

    readln;
  end.

