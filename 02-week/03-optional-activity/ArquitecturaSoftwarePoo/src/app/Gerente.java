package app;

public class Gerente extends Empleado {
    private double salarioBase;
    private double bonus;

    public Gerente(String id, String nombre, double salarioBase, double bonus) {
        super(id, nombre);
        this.salarioBase = salarioBase;
        this.bonus = bonus;
    }

    @Override
    public double calcularSalario() {
        return salarioBase + bonus;
    }
}