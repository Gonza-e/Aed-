Program Laboratorio13; 

{Un profesor necesita calcular el porcentaje de estudiantes que aprobaron un examen.
El archivo de texto "examenes.txt" contiene las calificaciones de los estudiantes (números enteros) 
en una sola línea separadas por espacios. Crea un programa en Pascal que lea estas calificaciones,
considere que cualquier calificación mayor o igual a 60 es aprobatoria, y calcule el porcentaje de estudiantes que aprobaron.

PREGUNTA: ¿Cuál es el porcentaje de estudiantes que aprobaron el examen?}

uses 

    sysutils;
    
var 
    sec: TextFile;
    v: Integer; 
    salida: Text;
    numero: String[2];
    porc: Real;
    cant_aprobados, cant_total: Integer; 
begin
    Assign(sec,'C:\Users\Gonzalo\Desktop\Algoritmos\2024\Ejercicios de la guia\Pascal\Secuencias\ArchivosParaUso\examenes.txt');
    Reset(sec); 
    Assign(salida,'numerosparaexcel.csv');
    Rewrite(salida);
    
    v:= 0;
    porc:=0; 
    cant_aprobados:=0;
    cant_total:=0;

    while not Eof(sec) do
        begin 
            if v >= 60 then
            begin
                cant_aprobados:= cant_aprobados + 1;
            end;
            cant_total:= cant_total + 1;
            Read(sec,v);
            Write(salida,v);
            Write(salida,';');

            if (cant_total = 20) or (cant_total = 40) or (cant_total = 60) or (cant_total = 80) or (cant_total = 100) then 
            begin
              WriteLn(salida);
            end;
            //Write(salida,';');   
        end; 



        WriteLn('Aprobados: ',cant_aprobados);
        WriteLn('total: ',cant_total);
        Close(sec);
        porc:= (cant_aprobados/cant_total)*100;
        WriteLn('el porcentaje de alumnos aprobados es: ',porc:0:2,'%');
        ReadLn;
end.

