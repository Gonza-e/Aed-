program Laboratorio14;


uses
  SysUtils;
type
 FORM = record
  legajo: integer;
  apellido_nombre: String[30];
  anio_ingreso: integer;
  carrera: String[30];
 end;
var
  arch: File of FORM;
  reg: FORM;
  total_alumnos: integer;
  prom: real;
begin
     Assign(arch,'C:\Users\Gonzalo\Desktop\Algoritmos\2024\Ejercicios de la guia\Pascal\TratamientoArchivos\Archivos\LD04.dat');
     reset(arch);
     total_alumnos:=0; prom:=0;

     while not Eof(arch) do
     begin
       total_alumnos:= total_alumnos + 1;
       Read(arch,reg);
     end;

     close(arch);
     prom:= (total_alumnos/4);
     writeln('Total de alumnos: ',total_alumnos);
     writeln('El promedio de alumnos por carrera es de: ',prom:0:2);
     Readln;
end.
