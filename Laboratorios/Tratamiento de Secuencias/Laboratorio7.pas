Program Laboratorio7; 

{Se tiene la información de ventas en un archivo de caracteres "LD03.txt" con el siguiente formato:
nro_cliente(2dig) tipo_cliente (JoF) ventas_realizadas (1dig) importes(3digitos)
el tipo de cliente puede ser F (Físico) o J (Juridico)
Tendremos tantos importes como ventas realizadas.

Por ejemplo:
01F322558623702J285162303F5523121222555127

El cliente 01 es del tipo Físico y compro 3 artículos, cada uno por un importe de 225,586 y 237 (esto es considerada como una venta)
El cliente 02 es del tipo Jurídico y compro 2 artículos, cada uno por un importe de 851, 623
El cliente 04 es del tipo Físico y compró 5 artículos, cada uno por un importe de 523 121 222 555 127
En el ejemplo se hicieron 3 ventas en total, 2 a clientes del tipo Jurídico y 1 clientes del tipo Físico

PREGUNTA: ¿Cuál es el importe promedio de las ventas a clientes del tipo físico?}

uses 
 math,
 sysutils; 
var 
 sec: TextFile;
 v: Char; 
 importe, k, ventas, i: Integer;
 importe_temp: LongInt;
 prom: Real; 
 fisico: Boolean; 

 procedure inicializar(); 
 begin
   importe:= 0; importe_temp:= 0; k:= 0; ventas:= 0; prom:= 0; 
   fisico:= False;
 end;
begin
  Assign(sec,'C:\Users\Gonzalo\Desktop\Algoritmos\2024\Ejercicios de la guia\Pascal\Secuencias\ArchivosParaUso\LD03.txt');
  Reset(sec); 
  Read(sec,v); 
  inicializar();

  while not Eof(sec) do
  begin
    Read(sec,v);
    Read(sec,v);


    if (v = 'F') then
    begin
      fisico:= true;
    end;
    Read(sec,v);

    if fisico then
    begin
      k:= strtoint(v);
      ventas:= ventas + k; 
      k:= k * 3; 

      Read(sec,v); 

      while k > 0 do
      begin
        for i:= 3 downto 1 do 
        begin
          importe_temp:= importe_temp + (strtoint(v) * (10**i-1));
          Read(sec,v);
          k:= k-1;
        end; 
        importe:= importe + importe_temp; 
        importe_temp:= 0;
      end;
      fisico:= False;
    end;
    
  end;
  Close(sec);
  prom:= (importe/ventas);
  WriteLn('El importe promedio es de: ',prom:0:1);
  WriteLn('Importe: ',importe,'$');
  WriteLn('Ventas: ',ventas);
end.