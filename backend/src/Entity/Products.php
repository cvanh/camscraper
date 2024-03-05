<?php

namespace App\Entity;

use ApiPlatform\Metadata\ApiResource;
use App\Repository\ProductsRepository;
use Doctrine\DBAL\Types\Types;
use Doctrine\ORM\Mapping as ORM;

#[ORM\Entity(repositoryClass: ProductsRepository::class)]
#[ApiResource]
class Products
{
    #[ORM\Id]
    #[ORM\GeneratedValue]
    #[ORM\Column]
    private ?int $id = null;

    #[ORM\Column(length: 255)]
    private ?string $mpn = null;

    #[ORM\Column(type: Types::TEXT)]
    private ?string $name = null;

    #[ORM\Column(type: Types::TEXT)]
    private ?string $source = null;

    #[ORM\Column(type: Types::DATETIME_MUTABLE)]
    private ?\DateTimeInterface $scrape_time = null;

    public function getId(): ?int
    {
        return $this->id;
    }

    public function getMpn(): ?string
    {
        return $this->mpn;
    }

    public function setMpn(string $mpn): static
    {
        $this->mpn = $mpn;

        return $this;
    }

    public function getName(): ?string
    {
        return $this->name;
    }

    public function setName(string $name): static
    {
        $this->name = $name;

        return $this;
    }

    public function getSource(): ?string
    {
        return $this->source;
    }

    public function setSource(string $source): static
    {
        $this->source = $source;

        return $this;
    }

    public function getScrapeTime(): ?\DateTimeInterface
    {
        return $this->scrape_time;
    }

    public function setScrapeTime(\DateTimeInterface $scrape_time): static
    {
        $this->scrape_time = $scrape_time;

        return $this;
    }
}
