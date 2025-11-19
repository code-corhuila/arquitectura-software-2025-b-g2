package app;

public class Administrativo extends Empleado {
    private double salarioMensual;

    public Administrativo(String id, String nombre, double salarioMensual) {
        super(id, nombre);
        this.salarioMensual = salarioMensual;
    }

    @Override
    public double calcularSalario() {
        return salarioMensual;
    }
}