package maybe

import "database/sql/driver"

func (o *Maybe[T]) Scan(src any) error {
	if src == nil {
		*o = None[T]()
		return nil
	}

	var v T

	if err := sqlConvertAssign(&v, src); err != nil {
		return err
	}

	*o = Some(v)
	
	return nil
}

func (o Maybe[T]) Value() (driver.Value, error) {
	if v, ok := o.Get(); ok {
		return driver.DefaultParameterConverter.ConvertValue(v)
	}
	
	return nil, nil
}
