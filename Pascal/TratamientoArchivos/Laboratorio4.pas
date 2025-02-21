Program Laboratorio4; 

{* Se tiene un archivo que almacena datos de jugadoras de FIFA a nivel mundial
(entrada-lab-ej4.csv) y contiene la siguiente información: ID de la jugadora,
versión de FIFA, nombre de la jugadora, edad, altura, peso, país de nacimiento
y pie preferido.

Se pide: ¿cuál es el porcentaje de jugadoras de treinta años o más que prefieren
utilizar el pie derecho para patear? *}

uses 
 csvdocument,
 sysutils;
type 
 JUGADORAS = record
   id: Integer; 
   version: String[5]; 
   name: String[50];
   edad: Integer;
   altura: Integer; 
   peso: Integer;
   nacionalidad: String[40]; 
   pie_favorito: String[20];
  end; 
var 
 arch: TCSVDocument;
 reg: JUGADORAS; 
 cant_jugadoras, pie_derecho, i: Integer; 
 porc: Real; 

 procedure inicializar(); 
 begin
   arch:= TCSVDocument.create();
   cant_jugadoras:= 0; pie_derecho:= 0;
   porc:= 0;
 end;

begin
    inicializar();
 
    arch.delimiter:= ';';
    arch.LoadFromFile('C:\Users\Gonzalo\Desktop\Algoritmos\2024\Ejercicios de la guia\Pascal\TratamientoArchivos\Archivos\laboratorio4.csv');
  
    for i:= 1 to (arch.rowcount - 1) do 
    begin
      cant_jugadoras:= cant_jugadoras + 1; 

      reg.id:= StrToInt(arch.cells[0,i]);
      reg.version:= arch.cells[1,i];
      reg.name:= arch.cells[2,i];
      reg.edad:= StrToInt(arch.cells[3,i]);
      reg.altura:= StrToInt(arch.cells[4,i]);
      reg.peso:= StrToInt(arch.cells[5,i]);
      reg.nacionalidad:= arch.cells[6,i];
      reg.pie_favorito:= arch.cells[7,i];

      if (reg.edad >= 30) and (reg.pie_favorito = 'Right') then
      begin
        pie_derecho:= pie_derecho + 1;
      end;
    end;
    porc:= (pie_derecho/cant_jugadoras) * 100;
    WriteLn('El porcentaje de jugadoras que cumplen con las condiciones es de: ',porc:0:2,'%');

    {WriteLn(cant_jugadoras);
    WriteLn(pie_derecho);}
 
    arch.free;
    ReadLn;
end.