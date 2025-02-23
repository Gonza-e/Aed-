program ejercicio3;

var
    cont: Integer;
    contgral: Integer;
    porcentaje: Real;
    calificacion: Integer;
    sec: Text;

begin
    Assign(sec, 'C:\Users\Gonzalo\Desktop\Algoritmos\2024\Ejercicios de la guia\Pascal\Secuencias\ArchivosParaUso\examenes.txt');
    {$I-}
    Reset(sec);
    {$I+}
    Read(sec, calificacion);

    cont := 0;
    contgral := 0;
    
    while not Eof(sec) do
    begin
        if calificacion >= 60 then
            begin 
                cont := cont + 1;
            end;
        contgral := contgral + 1;
        Read(sec, calificacion);
        //WriteLn(calificacion);
        //WriteLn('General: ',contgral);
    end;

    if calificacion >= 60 then  //para el ultimo numero 
        begin
            cont := cont + 1;
        end;
    contgral := contgral + 1;

    Close(sec);
    
    porcentaje := (cont / contgral) * 100;
    readln;
    //WriteLn('El porcentaje de aprobados es de: ', porcentaje:0:2);
end.
