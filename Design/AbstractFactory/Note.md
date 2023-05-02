# 抽象工厂
> 代码例子中具体产品是通过实现抽象产品接口实现的，还有一种方法是定义一个抽象产品struct去实现这个抽象接口，具体产品去继承这个struct。  
> 在这种实现方式中，我们通常会定义一个抽象产品 AbstractProduct，它是一个包含了一些共同特性的结构体，如产品名称、价格等等，并实现了 Product 接口。然后具体产品可以通过继承 AbstractProduct 并实现 Product 接口来实现自己的特定功能。下面是一个简单的示例代码： 
## 抽象工厂举例
```go
// 定义抽象产品接口
type Product interface {
    GetName() string
}

// 定义抽象产品结构体
type AbstractProduct struct {
    name string
}

func (p *AbstractProduct) GetName() string {
    return p.name
}

// 定义具体产品A
type ConcreteProductA struct {
    AbstractProduct
}

// 定义具体产品B
type ConcreteProductB struct {
    AbstractProduct
}

```

## 异同
在代码实现上，工厂模式和抽象工厂的不同主要体现在两个方面：产品层次结构和工厂接口。

1. 产品层次结构  

    在工厂模式中，产品的层次结构比较简单，只有一个抽象产品类和多个具体产品类。这些具体产品类都实现了同一个抽象产品类的接口，表示它们是同一种类型的产品。  

    在抽象工厂中，产品的层次结构比较复杂，通常包含多个抽象产品类和多个具体产品类。这些抽象产品类都定义了一些通用的属性和方法，而具体产品类则实现了这些抽象产品类的接口，并增加了一些具体的属性和方法。

2. 工厂接口  

    在工厂模式中，只有一个工厂接口和多个具体工厂实现。这个工厂接口定义了一个创建产品的方法，具体工厂则实现这个方法，从而创建具体的产品对象。

    在抽象工厂中，有多个抽象工厂接口和多个具体工厂实现。每个抽象工厂接口都定义了一组创建产品的方法，而具体工厂则实现这些方法，从而创建多个相互关联的产品族。

下面是一个简单的示例代码，展示了工厂模式和抽象工厂的区别：
```go
// 工厂模式示例代码
package factory

// 抽象产品类
type Product interface {
    Name() string
}

// 具体产品类
type ConcreteProductA struct{}

func (p *ConcreteProductA) Name() string {
    return "Product A"
}

type ConcreteProductB struct{}

func (p *ConcreteProductB) Name() string {
    return "Product B"
}

// 工厂接口
type Factory interface {
    CreateProduct() Product
}

// 具体工厂类
type ConcreteFactory struct{}

func (f *ConcreteFactory) CreateProduct() Product {
    return &ConcreteProductA{}
}

// 抽象工厂示例代码
package abstractfactory

// 抽象产品类
type Phone interface {
    Name() string
    OS() string
}

type Tablet interface {
    Name() string
    OS() string
}

// 具体产品类
type IPhone struct{}

func (p *IPhone) Name() string {
    return "iPhone"
}

func (p *IPhone) OS() string {
    return "iOS"
}

type IPad struct{}

func (p *IPad) Name() string {
    return "iPad"
}

func (p *IPad) OS() string {
    return "iOS"
}

// 抽象工厂接口
type Factory interface {
    CreatePhone() Phone
    CreateTablet() Tablet
}

// 具体工厂类
type AppleFactory struct{}

func (f *AppleFactory) CreatePhone() Phone {
    return &IPhone{}
}

func (f *AppleFactory) CreateTablet() Tablet {
    return &IPad{}
}
```