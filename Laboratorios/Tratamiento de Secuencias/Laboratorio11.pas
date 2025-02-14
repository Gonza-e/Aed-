Program Laboratorio11;

{Crear un programa en Pascal que lea una lista de edades (números enteros) 
desde un archivo de texto llamado problema31.txt, calcule cuantas personas de la categoría "niños y jóvenes" 
(de 0 a 14 años) hay, y el porcentaje de "niños y jóvenes" sobre el total de personas. El programa debe mostrar 
el porcentaje de "niños y jóvenes" sobre el total de personas.

PREGUNTA: ¿Cuál es el porcentaje de edades que corresponden a "niños y jóvenes❞ sobre el total de edades del archivo dado?}

uses
  sysutils;
var
  sec: TextFile;
  v: Integer; 
  edad_temp, ninyjovenes, personas_total: Integer;
  porc: Real; 
begin
  Assign(sec,'C:\Users\Gonzalo\Desktop\Algoritmos\2024\Ejercicios de la guia\Pascal\Secuencias\ArchivosParaUso\problema31.txt');
  Reset(sec); 
  edad_temp:= 0; ninyjovenes:= 0; personas_total:= 0; porc:= 0;

  while not Eof(sec) do
  begin 

    if (v > 0) and (v < 14) then 
    begin
      ninyjovenes:= ninyjovenes + 1; 
    end;
    personas_total:= personas_total + 1;
    Read(sec,v);
    
  end; 


  if (v > 0) and (v < 14) then 
  begin
    ninyjovenes:= ninyjovenes + 1;
    personas_total:= personas_total + 1;  
  end
  else 
  begin
    personas_total:= personas_total + 1; 
  end;
  
  Close(sec);
  {WriteLn(ninyjovenes);
  WriteLn(personas_total);}
  porc:= (ninyjovenes/personas_total)*100;
  WriteLn('El porcentaje de ninios y jovenes sobre el total de personas es de: ',porc:0:2,'%');
  ReadLn;
end. 
