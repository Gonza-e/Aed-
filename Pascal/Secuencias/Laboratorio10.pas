Program Laboratorio10; 

{Un profesor necesita calcular el porcentaje de estudiantes que aprobaron un examen.
El archivo de texto "examenes.txt" contiene las calificaciones de los estudiantes (números enteros) 
en una sola línea separadas por espacios. Crea un programa en Pascal que lea estas calificaciones,
considere que cualquier calificación mayor o igual a 60 es aprobatoria, y calcule el porcentaje de estudiantes que aprobaron.

PREGUNTA: ¿Cuál es el porcentaje de estudiantes que aprobaron el examen?}

uses 
 sysutils;
var 
  sec: TextFile;
  v: Char; 
  porc: Real; 
  cant_aprobados, cant_total, calificacion: Integer;
  function conv_entero(c:Char):Integer;
  begin
    case c of 
     '0': conv_entero:= 0;
     '1': conv_entero:= 1;
     '2': conv_entero:= 2;
     '3': conv_entero:= 3;
     '4': conv_entero:= 4;
     '5': conv_entero:= 5;
     '6': conv_entero:= 6;
     '7': conv_entero:= 7;
     '8': conv_entero:= 8;
     '9': conv_entero:= 9;
    end;
  end; 
begin
  Assign(sec,'C:\Users\Gonzalo\Desktop\Algoritmos\2024\Ejercicios de la guia\Pascal\Secuencias\ArchivosParaUso\examenes.txt');
  Reset(sec); 
  Read(sec,v); 
  cant_aprobados:= 0; cant_total:= 0; calificacion:= 0; porc:= 0; 

  while not Eof(sec) do
  begin
    while (not Eof(sec)) and (v <> ' ') do
    begin
      calificacion:= conv_entero(v)*10;
      Read(sec,v);
      calificacion:= calificacion + conv_entero(v);
      Read(sec,v); 
    end;
    if calificacion > 60 then 
    begin
      cant_aprobados+=1;
    end;
    cant_total+=1;
    Read(sec,v);
  end; 
  WriteLn(cant_aprobados);
  WriteLn(cant_total);
  Close(sec);
  porc:= (cant_aprobados/cant_total)*100;
  WriteLn('El porcentaje de alumnos aprobados es de: ',porc:0:2,'%');
end. 