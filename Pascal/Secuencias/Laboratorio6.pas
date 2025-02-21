Program Laboratorio6; 

{Crear un programa en Pascal que lea una lista de edades (números enteros) desde un
archivo de texto llamado problema31.txt, calcule cuantos adultos (entre 15 y 64 años) hay 
y el promedio de edad de los mismos. El programa debe mostrar el promedio de edad de los adultos.

PREGUNTA: ¿Cuál es el promedio de edad de los adultos del archivo dado?}

uses
 sysutils;
var 
 sec: TextFile;
 v: String; 
 edad, edad_temp, cant_adultos: Integer;
 prom: Real; 
begin
  Assign(sec,'C:\Users\Gonzalo\Desktop\Algoritmos\2024\Ejercicios de la guia\Pascal\Secuencias\ArchivosParaUso\problema31.txt');
  Reset(sec); 
  ReadLn(sec,v); 
  edad:= 0; cant_adultos:= 0; prom:= 0; edad_temp:= 0;
  while not Eof(sec) do 
  begin

    edad_temp:= strtoint(v);
    
    if (edad_temp > 15) and (edad_temp < 64) then 
    begin
      cant_adultos+=1; 
      edad:= edad + edad_temp;
    end;
    ReadLn(sec,v);
  end;
  Close(sec);
  prom:= (edad/cant_adultos);
  WriteLn('El promedio de edad de adultos de entre 15 y 64 años es de: ',prom:0:1);
  {WriteLn(edad);
  WriteLn(cant_adultos);}
end.
