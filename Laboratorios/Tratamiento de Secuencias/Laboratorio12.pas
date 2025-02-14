Program Laboratorio12; 

{Un Ministerio de la Provincia del Chaco necesita procesar datos del último Censo Nacional de Población, 
Hogares y Viviendas 2022 (Censo 2022) realizado por el INDEC. Para ello dispone de un archivo de texto 
censo_phv.txt que contiene información de la Provincia del Chaco: Nro_Vivienda 
(cantidad de caracteres indeterminada, termina con #), Tipo_vivienda ('P'-Particular o 'C'-Colectiva) 
y Condicion_Vivienda (Y-Muy Buena, 'B'- Buena o 'M'-Mala).

Pregunta: ¿Qué porcentaje de viviendas "particulares" de condición "muy buena" existen, sobre el total de viviendas "particulares"?}

uses 
 sysutils; 
var 
 sec: TextFile;
 v: Char; 
 particular, muybuena: Boolean;
 viviendas_total, viviendas_cumplen: Integer;
 porc: Real; 

 procedure inicializar();
 begin
   particular:= False; muybuena:= False; 
   viviendas_total:= 0; viviendas_cumplen:= 0; porc:= 0; 
 end;
begin
  Assign(sec,'C:\Users\Gonzalo\Desktop\Algoritmos\2024\Ejercicios de la guia\Pascal\Secuencias\ArchivosParaUso\censo_phv.txt'); 
  Reset(sec); 
  Read(sec,v);
  inicializar(); 
  
  while (not Eof(sec))  do
  begin

    while (not Eof(sec)) and (v <> '#') do
    begin
      Read(sec,v);
    end;
    Read(sec,v); 
    
    if v = 'P' then
    begin
      particular:= True;
      viviendas_total:= viviendas_total + 1;
    end;
    Read(sec,v); 

    if v = 'Y' then 
    begin
      muybuena:= True;
    end;

    if (muybuena) and (particular) then 
    begin
      viviendas_cumplen:= viviendas_cumplen + 1;
      muybuena:= False; particular:= False;
    end;
  end;
  Close(sec);
  porc:= (viviendas_cumplen/viviendas_total)*100;
  WriteLn('El porcentaje de viviendas que cumplen con las condiciones es de: ',porc:0:2,'%');
  {WriteLn(viviendas_total);
  WriteLn(viviendas_cumplen);}
end. 