from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Status(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    Ok: _ClassVar[Status]
    Err: _ClassVar[Status]
Ok: Status
Err: Status

class Image(_message.Message):
    __slots__ = ("value",)
    VALUE_FIELD_NUMBER: _ClassVar[int]
    value: _containers.RepeatedScalarFieldContainer[int]
    def __init__(self, value: _Optional[_Iterable[int]] = ...) -> None: ...

class ProduceRequest(_message.Message):
    __slots__ = ("img",)
    IMG_FIELD_NUMBER: _ClassVar[int]
    img: Image
    def __init__(self, img: _Optional[_Union[Image, _Mapping]] = ...) -> None: ...

class ProduceResponse(_message.Message):
    __slots__ = ("status", "image")
    STATUS_FIELD_NUMBER: _ClassVar[int]
    IMAGE_FIELD_NUMBER: _ClassVar[int]
    status: Status
    image: str
    def __init__(self, status: _Optional[_Union[Status, str]] = ..., image: _Optional[str] = ...) -> None: ...
